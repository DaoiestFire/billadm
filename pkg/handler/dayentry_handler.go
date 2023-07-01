package handler

import (
	"fmt"

	"path"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/operation"
	"ljw/billadm/pkg/types"
	"ljw/billadm/utils/print"
	"ljw/billadm/utils/time"
)

var _ IResource = &DayEntryHandler{}

type DayEntryHandler struct {
	path     string
	dayEntry *types.DayEntry
}

func (dh *DayEntryHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case operation.Get:
		err = dh.get(resourceName, resources, cm, options)
	case operation.Delete:
		err = dh.delete(resourceName, resources, cm, options)
	case operation.Create:
		err = dh.create(resourceName, resources, cm, options)
	case operation.Modify:
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

}

func (dh *DayEntryHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (dh *DayEntryHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
