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
	"ljw/billadm/utils/time"
)

var _ IResource = &RecordHandler{}

type RecordHandler struct {
}

func (rh *RecordHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case constant.Get:
		err = rh.get(resourceName, resources, cm, options)
	case constant.Delete:
		err = rh.delete(resourceName, resources, cm, options)
	case constant.Create:
		err = rh.create(resourceName, resources, cm, options)
	case constant.Modify:
		err = rh.modify(resourceName, resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for RecordHandler", op)
	}
	return err
}

func (rh *RecordHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// get record 不需要了。直接get de获取全部record
	return nil
}

func (rh *RecordHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 删除一个record，先判断de是否存在
	// 必须要激活一个bill
	if cm.Config.CurrentBillName == "" {
		return fmt.Errorf("please activate a bill first")
	}

	year, month, _ := time.GetYearMonthDay(options.Time)
	fileName := fmt.Sprintf("%s.json", options.Time)
	dayEntryPath := path.Join(cm.Config.BillDataDir, cm.Config.CurrentBillName, year, month, fileName)

	// 如果文件不存在，就跳过
	if !fileutils.Exist(dayEntryPath) {
		logger.Infof("day entry [%s] doesn't existed", fileName)
		return nil
	}

	de, err := types.ReadOneDayEntry(dayEntryPath)
	if err != nil {
		return fmt.Errorf("read day entry [%s] failed ---> <%v>", fileName, err)
	}

	if err := de.DeleteRecord(options.Id); err != nil {
		return fmt.Errorf("delete record [id -> %03d] from day entry [%s] failed ---> <%s>", options.Id, fileName, err)
	}

	if de.Len() == 0 {
		err = resources.dayEntryHandler.Run(constant.Delete, resourceName, resources, cm, options)
		if err != nil {
			return fmt.Errorf("error after delete record ---> <%v>", err)
		}
		return nil
	}

	err = types.SaveOneDayEntry(dayEntryPath, de)
	if err != nil {
		return fmt.Errorf("save day entry [%s] failed after delete record [id -> %03d] ---> <%v>", fileName, options.Id, err)
	}
	return nil
}

func (rh *RecordHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 必须要激活一个bill
	if cm.Config.CurrentBillName == "" {
		return fmt.Errorf("please activate a bill first")
	}

	year, month, _ := time.GetYearMonthDay(options.Time)
	fileName := fmt.Sprintf("%s.json", options.Time)
	dayEntryPath := path.Join(cm.Config.BillDataDir, cm.Config.CurrentBillName, year, month, fileName)

	// 如果文件不存在，就创建de
	if !fileutils.Exist(dayEntryPath) {
		err := resources.dayEntryHandler.Run(constant.Create, resourceName, resources, cm, options)
		if err != nil {
			return fmt.Errorf("create day entry failed before create record ---> <%v>", err)
		}
	}

	de, err := types.ReadOneDayEntry(dayEntryPath)
	if err != nil {
		return fmt.Errorf("read day entry [%s] failed ---> <%v>", fileName, err)
	}

	r := de.AddRecord()
	if options.Cost == 0 {
		return fmt.Errorf("cost should not be zero when create record")
	}
	if options.Label > len(types.LabelList) {
		return fmt.Errorf("label should not be lager than [%d]", len(types.LabelList))
	}
	r.SetCost(options.Cost)
	// 描述可以是空的
	r.SetDescription(options.Description)
	r.SetLabels(types.LabelList[options.Label])

	err = types.SaveOneDayEntry(dayEntryPath, de)
	if err != nil {
		return fmt.Errorf("save day entry [%s] failed after delete record [id -> %03d] ---> <%v>", fileName, options.Id, err)
	}
	return nil
}

func (rh *RecordHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 必须要激活一个bill
	if cm.Config.CurrentBillName == "" {
		return fmt.Errorf("please activate a bill first")
	}

	year, month, _ := time.GetYearMonthDay(options.Time)
	fileName := fmt.Sprintf("%s.json", options.Time)
	dayEntryPath := path.Join(cm.Config.BillDataDir, cm.Config.CurrentBillName, year, month, fileName)

	// 如果文件不存在，就报错。无法修改一个不存在的de
	if !fileutils.Exist(dayEntryPath) {
		return fmt.Errorf("modify record failed ---> <day entry [%s] not exist>", fileName)
	}

	de, err := types.ReadOneDayEntry(dayEntryPath)
	if err != nil {
		return fmt.Errorf("read day entry [%s] failed ---> <%v>", fileName, err)
	}

	r, err := de.GetRecordById(options.Id)
	if err != nil {
		return err
	}
	if options.Cost != 0 {
		r.SetCost(options.Cost)
	}
	if options.Description == "" {
		r.SetDescription(options.Description)
	}
	if options.Label > len(types.LabelList) {
		return fmt.Errorf("label should not be lager than [%d]", len(types.LabelList))
	}
	if options.Label != 0 {
		r.SetLabels(types.LabelList[options.Label])
	}
	// 描述可以是空的
	err = types.SaveOneDayEntry(dayEntryPath, de)
	if err != nil {
		return fmt.Errorf("save day entry [%s] failed after delete record [id -> %03d] ---> <%v>", fileName, options.Id, err)
	}
	return nil
}
