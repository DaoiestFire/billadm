package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"k8s.io/klog/v2"

	constant "ljw/billadm/const"
	"ljw/billadm/pkg/api/service"
	"ljw/billadm/pkg/api/v1"
	"ljw/billadm/utils/fileutils"
	timeutils "ljw/billadm/utils/time"
)

var _ service.StorageServiceServer = &Storage{}

var storage *Storage
var once sync.Once

func NewStorage(installPath string) (*Storage, error) {
	if storage != nil {
		return storage, nil
	}
	var err error
	once.Do(func() {
		billadmDataDir := filepath.Join(installPath, constant.Data)
		storage = &Storage{
			billadmDataDir: billadmDataDir,
			billMapper:     make(map[string]*billManager),
		}
		err = storage.Initializer()
		currentBillNamePath := filepath.Join(billadmDataDir, constant.CurrentBillName)
		if fileutils.Exist(currentBillNamePath) {
			res, errTmp := fileutils.ReadFileString(currentBillNamePath)
			if errTmp != nil {
				err = errTmp
				return
			}
			res = strings.TrimSpace(res)
			_, err = storage.SetCurrentBillName(context.Background(), &service.SetCurrentBillNameRequest{Name: res})
		}
	})
	if err != nil {
		return nil, err
	}
	return storage, nil
}

// 考虑并发读写问题，Storage层设置读写锁
// bill层设置互斥锁
// 修改Storage的成员需要取得Storage的写锁，修改bill的成员需要先取得storage的读锁，再取得bill的互斥锁
// 硬盘操作较多，简化逻辑

type Storage struct {
	service.UnimplementedStorageServiceServer
	billadmDataDir string
	// currentBillName的bill必然是存在的
	currentBillName string
	billMapper      map[string]*billManager
	rwMutex         sync.RWMutex
}

func (s *Storage) GetCurrentBillName(ctx context.Context, request *service.GetCurrentBillNameRequest) (*service.GetCurrentBillNameResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Info("get current bill name")
	out := new(service.GetCurrentBillNameResponse)
	out.Name = s.currentBillName
	return out, nil
}

func (s *Storage) SetCurrentBillName(ctx context.Context, request *service.SetCurrentBillNameRequest) (*service.SetCurrentBillNameResponse, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	klog.Infof("set currrent bill name to : %s", request.Name)
	if request.Name == "" {
		s.currentBillName = ""
		return nil, nil
	}
	if _, ok := s.billMapper[request.Name]; !ok {
		return nil, fmt.Errorf("bill [%s] not found", request.Name)
	}
	s.currentBillName = request.Name
	return nil, nil
}

func (s *Storage) ListAllBill(ctx context.Context, request *service.ListAllBillRequest) (*service.ListAllBillResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Info("list all bills")
	out := new(service.ListAllBillResponse)
	billInfoList := make([]*service.BillInfo, 0, len(s.billMapper))
	for key := range s.billMapper {
		billInfoList = append(billInfoList, s.billMapper[key].bill.ToBillInfo())
	}
	out.BillList = billInfoList
	return out, nil
}

func (s *Storage) GetBill(ctx context.Context, request *service.GetBillRequest) (*service.GetBillResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Infof("get bill [%s]", request.Name)
	if _, ok := s.billMapper[request.Name]; !ok {
		return nil, fmt.Errorf("bill [%s] not found", request.Name)
	}
	out := new(service.GetBillResponse)
	out.Bill = s.billMapper[request.Name].bill.ToBillInfo()
	return out, nil
}

func (s *Storage) CreateBill(ctx context.Context, request *service.CreateBillRequest) (*service.CreateBillResponse, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	klog.Infof("create bill [%s]", request.Name)
	name := request.Name
	// 如果没有被标记未删除且bill存在，则报错
	if _, ok := s.billMapper[name]; ok {
		klog.Errorf("create bill failed: bill [%s] existed", request.Name)
		return nil, fmt.Errorf("bill [%s] existed", name)
	}
	// 创建一个新的bill
	s.billMapper[name] = &billManager{
		dataPath:      filepath.Join(s.billadmDataDir, name),
		billName:      name,
		dayEntryCache: make(map[string]v1.IDayEntry),
	}
	s.billMapper[name].bill = v1.NewBill(name)
	if err := fileutils.CreateDirectory(s.billMapper[name].dataPath); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Storage) DeleteBill(ctx context.Context, request *service.DeleteBillRequest) (*service.DeleteBillResponse, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	klog.Infof("delete bill [%s]", request.Name)
	name := request.Name
	if _, ok := s.billMapper[name]; !ok {
		klog.Errorf("delete bill failed: bill [%s] not found", request.Name)
		return nil, fmt.Errorf("bill [%s] not found", name)
	}
	if err := fileutils.RemoveDirectoryOrFile(s.billMapper[name].dataPath); err != nil {
		return nil, err
	}
	delete(s.billMapper, name)
	if strings.EqualFold(name, s.currentBillName) {
		s.currentBillName = ""
	}
	return nil, nil
}

func (s *Storage) ListAllDayEntry(ctx context.Context, request *service.ListAllDayEntryRequest) (*service.ListAllDayEntryResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Info("list all day entry")
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	out := new(service.ListAllDayEntryResponse)
	res := s.billMapper[s.currentBillName].ListAllDayEntry()
	dayEntryInfoList := make([]*service.DayEntryInfo, 0, len(res))
	for i := range res {
		dayEntryInfoList = append(dayEntryInfoList, res[i].ToDayEntryInfo())
	}
	out.DayEntryList = dayEntryInfoList
	return out, nil
}

func (s *Storage) GetDayEntry(ctx context.Context, request *service.GetDayEntryRequest) (*service.GetDayEntryResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Infof("get day entry [%s]", request.Name)
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	out := new(service.GetDayEntryResponse)
	dayEntry, err := s.billMapper[s.currentBillName].getDayEntry(request.Name)
	if err != nil {
		return nil, err
	}
	out.DayEntry = dayEntry.ToDayEntryInfo()
	return out, nil
}

func (s *Storage) CreateDayEntry(ctx context.Context, request *service.CreateDayEntryRequest) (*service.CreateDayEntryResponse, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	klog.Infof("create day entry [%s]", request.DayEntry.ObjectMeta.Name)
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	return nil, s.billMapper[s.currentBillName].createDayEntry(request.DayEntry.ObjectMeta.Name)
}

func (s *Storage) DeleteDayEntry(ctx context.Context, request *service.DeleteDayEntryRequest) (*service.DeleteDayEntryResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Infof("delete day entry [%s]", request.Name)
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	return nil, s.billMapper[s.currentBillName].deleteDayEntry(request.Name)
}

func (s *Storage) GetRecord(ctx context.Context, request *service.GetRecordRequest) (*service.GetRecordResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Infof("get record [%s/%s]", request.DayEntryName, request.Id)
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	out := new(service.GetRecordResponse)
	record, err := s.billMapper[s.currentBillName].getRecord(request.DayEntryName, request.Id)
	if err != nil {
		return nil, err
	}
	out.Record = record.ToRecordInfo()
	return out, nil
}

func (s *Storage) CreateRecord(ctx context.Context, request *service.CreateRecordRequest) (*service.CreateRecordEntryResponse, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	klog.Infof("create record [%s]", request.Record.ObjectMeta.Name)
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	record := v1.NewRecord("")
	record.FromRecordInfo(request.Record)
	return nil, s.billMapper[s.currentBillName].createRecord(record)
}

func (s *Storage) DeleteRecord(ctx context.Context, request *service.DeleteRecordRequest) (*service.DeleteRecordResponse, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	klog.Infof("delete record [%s]", request.Id)
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	return nil, s.billMapper[s.currentBillName].deleteRecord(request.DayEntryName, request.Id)
}

// Initializer 初始化Storage
func (s *Storage) Initializer() error {
	// 遍历数据目录下的文件，获取其中的目录项
	dirs, err := os.ReadDir(s.billadmDataDir)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		// 不是目录不处理
		if !dir.IsDir() {
			continue
		}
		billName := dir.Name()
		s.billMapper[billName] = &billManager{
			dataPath:      filepath.Join(s.billadmDataDir, billName),
			billName:      billName,
			dayEntryCache: make(map[string]v1.IDayEntry),
		}
		if err := s.billMapper[billName].initializer(); err != nil {
			errMsg := fmt.Errorf("initialize billManager [%s] failed -> <%v>", billName, err)
			klog.Error(errMsg.Error())
			return errMsg
		}
	}
	return nil
}

// Finalizer 终止程序时执行，将内存中的数据刷回磁盘
func (s *Storage) Finalizer() error {
	// 遍历所有的依然存在的bill调用器finalizer
	for k := range s.billMapper {
		if err := s.billMapper[k].finalizer(); err != nil {
			return err
		}
	}
	currentBillNamePath := filepath.Join(s.billadmDataDir, constant.CurrentBillName)
	if err := fileutils.WriteFileString(currentBillNamePath, s.currentBillName); err != nil {
		return err
	}
	return nil
}

type billManager struct {
	dataPath      string
	billName      string
	bill          v1.IBill
	dayEntryCache map[string]v1.IDayEntry
}

func (b *billManager) getRecord(deName, id string) (v1.IRecord, error) {
	if _, ok := b.dayEntryCache[deName]; !ok {
		return nil, fmt.Errorf("DayEntry [%s] not found", deName)
	}
	// 查看de中是否存在record
	de := b.dayEntryCache[deName]
	return de.GetRecord(id)
}

func (b *billManager) getDayEntry(deName string) (v1.IDayEntry, error) {
	if _, ok := b.dayEntryCache[deName]; !ok {
		return nil, fmt.Errorf("DayEntry [%s] not found", deName)
	}
	return b.dayEntryCache[deName], nil
}

func (b *billManager) deleteDayEntry(deName string) error {
	if _, ok := b.dayEntryCache[deName]; !ok {
		return fmt.Errorf("DayEntry [%s] not found", deName)
	}
	// 删除内存中的de缓存
	delete(b.dayEntryCache, deName)
	// 删除磁盘中的de文件
	y, m, _ := timeutils.GetYearMonthDay(deName)
	dayEntryFilePath := filepath.Join(b.dataPath, y, m, deName+".json")
	if err := fileutils.RemoveFileRecursive(dayEntryFilePath, b.dataPath); err != nil {
		return err
	}
	return nil
}

func (b *billManager) deleteRecord(deName, id string) error {
	if _, ok := b.dayEntryCache[deName]; !ok {
		return fmt.Errorf("DayEntry [%s] not found", deName)
	}
	de := b.dayEntryCache[deName]
	return de.DeleteRecord(id)
}

func (b *billManager) createDayEntry(deName string) error {
	if _, ok := b.dayEntryCache[deName]; ok {
		return fmt.Errorf("DayEntry [%s] exsited", deName)
	}
	y, m, _ := timeutils.GetYearMonthDay(deName)
	b.dayEntryCache[deName] = v1.NewDayEntry(deName)
	if err := fileutils.CreateDirectory(filepath.Join(b.dataPath, y, m)); err != nil {
		return err
	}
	return nil
}

func (b *billManager) createRecord(record v1.IRecord) error {
	deName := record.GetDayEntryName()
	if _, ok := b.dayEntryCache[deName]; !ok {
		return fmt.Errorf("DayEntry [%s] not found", deName)
	}
	de := b.dayEntryCache[deName]
	de.AddRecord(record)
	return nil
}

func (b *billManager) ListAllDayEntry() []v1.IDayEntry {
	res := make([]v1.IDayEntry, 0, len(b.dayEntryCache))
	for key := range b.dayEntryCache {
		res = append(res, b.dayEntryCache[key])
	}
	return res
}

func (b *billManager) initializer() error {
	// 初始化bill配置文件
	billConfigPath := filepath.Join(b.dataPath, constant.BillConfig)
	data, err := fileutils.ReadFileByte(billConfigPath)
	if err != nil {
		return err
	}
	b.bill = &v1.Bill{}
	if err := b.bill.UnmarshalFrom(data); err != nil {
		return err
	}
	// 创建缓存
	allFilePath, err := fileutils.FindAllFileFromDirectory(b.dataPath)
	if err != nil {
		return err
	}
	reg := regexp.MustCompile("^(.*)[0-9]{4}-[0-9]{2}-[0-9]{2}\\.json$")
	for _, p := range allFilePath {
		if match := reg.MatchString(p); match {
			data, err := fileutils.ReadFileByte(p)
			if err != nil {
				return err
			}
			de := &v1.DayEntry{}
			if err := de.UnmarshalFrom(data); err != nil {
				return err
			}
			b.dayEntryCache[de.Name] = de
		}
	}
	return nil
}

// 终止程序时执行，将内存中的数据刷回磁盘。会有Storage的Finalizer调用
func (b *billManager) finalizer() error {
	for _, de := range b.dayEntryCache {
		y, m, _ := timeutils.GetYearMonthDay(de.GetName())
		targetPath := filepath.Join(b.dataPath, y, m)
		if !fileutils.Exist(targetPath) {
			if err := fileutils.CreateDirectory(targetPath); err != nil {
				return err
			}
		}
		data, err := de.MarshalTo()
		if err != nil {
			return err
		}
		if err := fileutils.WriteFileByte(filepath.Join(targetPath, de.GetName()+".json"), data); err != nil {
			return err
		}
	}
	targetPath := filepath.Join(b.dataPath, constant.BillConfig)
	data, err := b.bill.MarshalTo()
	if err != nil {
		return err
	}
	if err := fileutils.WriteFileByte(targetPath, data); err != nil {
		return err
	}
	return nil
}
