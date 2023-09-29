package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	logger "k8s.io/klog/v2"

	constant "ljw/billadm/const"
	"ljw/billadm/pkg/api/v1"
	configutils "ljw/billadm/utils/config"
	"ljw/billadm/utils/fileutils"
	timeutils "ljw/billadm/utils/time"
)

var storage *Storage
var once sync.Once

func GetStorage() (*Storage, error) {
	if storage != nil {
		return storage, nil
	}
	var err error
	once.Do(func() {
		billadmDataDir := filepath.Join(configutils.InstallPath, constant.Data)
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
			err = storage.SetCurrentBillName(res)
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
	billadmDataDir string
	// currentBillName的bill必然是存在的
	currentBillName string
	billMapper      map[string]*billManager
	rwMutex         sync.RWMutex
}

func (s *Storage) GetCurrentBillName() string {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	return s.currentBillName
}

func (s *Storage) SetCurrentBillName(name string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	if name == "" {
		s.currentBillName = ""
		return nil
	}
	if _, ok := s.billMapper[name]; !ok {
		return fmt.Errorf("bill [%s] not existed", name)
	}
	s.currentBillName = name
	return nil
}

func (s *Storage) ListAllBill() []v1.IBill {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	res := make([]v1.IBill, 0, len(s.billMapper))
	for key := range s.billMapper {
		res = append(res, s.billMapper[key].bill)
	}
	return res
}

func (s *Storage) ListAllDayEntry() ([]v1.IDayEntry, error) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	return s.billMapper[s.currentBillName].ListAllDayEntry(), nil
}

func (s *Storage) GetRecord(deName, id string) (v1.IRecord, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	return s.billMapper[s.currentBillName].getRecord(deName, id)
}

func (s *Storage) GetDayEntry(name string) (v1.IDayEntry, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	return s.billMapper[s.currentBillName].getDayEntry(name)
}

func (s *Storage) GetBill(name string) (v1.IBill, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if _, ok := s.billMapper[name]; !ok {
		return nil, fmt.Errorf("bill [%s] not existed", name)
	}
	return s.billMapper[name].bill, nil
}

func (s *Storage) DeleteRecord(deName, id string) error {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.currentBillName == "" {
		return fmt.Errorf("please activate a bill")
	}
	return s.billMapper[s.currentBillName].deleteRecord(deName, id)
}

func (s *Storage) DeleteDayEntry(name string) error {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.currentBillName == "" {
		return fmt.Errorf("please activate a bill")
	}
	return s.billMapper[s.currentBillName].deleteDayEntry(name)
}

func (s *Storage) DeleteBill(name string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	if _, ok := s.billMapper[name]; !ok {
		return fmt.Errorf("bill [%s] not existed", name)
	}
	if err := fileutils.RemoveDirectoryOrFile(s.billMapper[name].dataPath); err != nil {
		return err
	}
	delete(s.billMapper, name)
	if strings.EqualFold(name, s.currentBillName) {
		s.currentBillName = ""
	}
	return nil
}

func (s *Storage) CreateBill(name string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	// 如果没有被标记未删除且bill存在，则报错
	if _, ok := s.billMapper[name]; ok {
		return fmt.Errorf("bill [%s] existed", name)
	}
	// 创建一个新的bill
	s.billMapper[name] = &billManager{
		dataPath:      filepath.Join(s.billadmDataDir, name),
		billName:      name,
		dayEntryCache: make(map[string]*v1.DayEntry),
	}
	s.billMapper[name].bill = v1.NewBill(name)
	if err := fileutils.CreateDirectory(s.billMapper[name].dataPath); err != nil {
		return err
	}
	return nil
}

func (s *Storage) CreateDayEntry(name string) error {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.currentBillName == "" {
		return fmt.Errorf("please activate a bill")
	}
	return s.billMapper[s.currentBillName].createDayEntry(name)
}

func (s *Storage) CreateRecord(name string) (v1.IRecord, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.currentBillName == "" {
		return nil, fmt.Errorf("please activate a bill")
	}
	return s.billMapper[s.currentBillName].createRecord(name)
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
			dayEntryCache: make(map[string]*v1.DayEntry),
		}
		if err := s.billMapper[billName].initializer(); err != nil {
			errMsg := fmt.Errorf("initialize billManager [%s] failed -> <%v>", billName, err)
			logger.Error(errMsg.Error())
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
	bill          *v1.Bill
	dayEntryCache map[string]*v1.DayEntry
	mutex         sync.Mutex
}

func (b *billManager) getRecord(deName, id string) (v1.IRecord, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.dayEntryCache[deName]; !ok {
		return nil, fmt.Errorf("DayEntry [%s] not exsited", deName)
	}
	// 查看de中是否存在record
	de := b.dayEntryCache[deName]
	return de.GetRecord(id)
}

func (b *billManager) getDayEntry(deName string) (v1.IDayEntry, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.dayEntryCache[deName]; !ok {
		return nil, fmt.Errorf("DayEntry [%s] not exsited", deName)
	}
	return b.dayEntryCache[deName], nil
}

func (b *billManager) deleteDayEntry(deName string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.dayEntryCache[deName]; !ok {
		return fmt.Errorf("DayEntry [%s] not exsited", deName)
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
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.dayEntryCache[deName]; !ok {
		return fmt.Errorf("DayEntry [%s] not exsited", deName)
	}
	de := b.dayEntryCache[deName]
	return de.DeleteRecord(id)
}

func (b *billManager) createDayEntry(deName string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
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

func (b *billManager) createRecord(deName string) (v1.IRecord, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.dayEntryCache[deName]; !ok {
		return nil, fmt.Errorf("DayEntry [%s] not exsited", deName)
	}
	de := b.dayEntryCache[deName]
	return de.AddRecord(), nil
}

func (b *billManager) ListAllDayEntry() []v1.IDayEntry {
	b.mutex.Lock()
	defer b.mutex.Unlock()
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
		y, m, _ := timeutils.GetYearMonthDay(de.Name)
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
		if err := fileutils.WriteFileByte(filepath.Join(targetPath, de.Name+".json"), data); err != nil {
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
