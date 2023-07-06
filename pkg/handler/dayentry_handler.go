package handler

import (
	"fmt"

	"path"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/types"
	"ljw/billadm/utils/fileutils"
	"ljw/billadm/utils/logger"
	"ljw/billadm/utils/print"
	"ljw/billadm/utils/time"
)

var _ IResource = &DayEntryHandler{}

type DayEntryHandler struct {
}

func (dh *DayEntryHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case constant.Get:
		err = dh.get(resourceName, resources, cm, options)
	case constant.Delete:
		err = dh.delete(resourceName, resources, cm, options)
	case constant.Create:
		err = dh.create(resourceName, resources, cm, options)
	case constant.Modify:
		err = dh.modify(resourceName, resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for DayEntryHandler", op)
	}
	return err
}

// 打印Time时间的DE中的record
func (dh *DayEntryHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 必须要激活一个bill
	if cm.Config.CurrentBillName == "" {
		return fmt.Errorf("please activate a bill first")
	}

	year, month, _ := time.GetYearMonthDay(options.Time)
	fileName := fmt.Sprintf("%s.json", options.Time)
	dayEntryPath := path.Join(cm.Config.BillDataDir, cm.Config.CurrentBillName, year, month, fileName)
	de, err := types.ReadOneDayEntry(dayEntryPath)
	if err != nil {
		return fmt.Errorf("read day entry [%s] failed ---> <%v>", fileName, err)
	}

	fmt.Printf("All records of day entey [%s] :\n", fileName)
	for _, r := range de.GetRecords() {
		print.OneRecordPrint(r)
	}
	return nil
}

func (dh *DayEntryHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 必须要激活一个bill
	if cm.Config.CurrentBillName == "" {
		return fmt.Errorf("please activate a bill first")
	}

	year, month, _ := time.GetYearMonthDay(options.Time)
	fileName := fmt.Sprintf("%s.json", options.Time)
	dayEntryPath := path.Join(cm.Config.BillDataDir, cm.Config.CurrentBillName, year, month, fileName)

	// 直接删除对应的de
	err := fileutils.RemoveFileRecursive(dayEntryPath, cm.Config.BillDataDir)
	if err != nil {
		logger.Errorf("delete day entry failed ---> <%v>", err)
	}
	return nil
}

func (dh *DayEntryHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 必须要激活一个bill
	if cm.Config.CurrentBillName == "" {
		return fmt.Errorf("please activate a bill first")
	}

	year, month, _ := time.GetYearMonthDay(options.Time)
	fileName := fmt.Sprintf("%s.json", options.Time)
	dayEntryPath := path.Join(cm.Config.BillDataDir, cm.Config.CurrentBillName, year, month, fileName)

	// 如果文件存在，就跳过
	if fileutils.Exist(dayEntryPath) {
		logger.Infof("day entry [%s] existed", fileName)
	}

	// 不存在时需要先创建目录
	err := fileutils.CreateDir(path.Dir(dayEntryPath))
	if err != nil {
		return fmt.Errorf("create day entry directory failed ---> <%v>", err)
	}
	// 生成新的de,存储到本地
	de := types.NewDayEntry(options.Time)
	err = types.SaveOneDayEntry(dayEntryPath, de)
	if err != nil {
		return fmt.Errorf("create day entry [%s] failed --> <%v>", dayEntryPath, err)
	}
	return nil
}

func (dh *DayEntryHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}
