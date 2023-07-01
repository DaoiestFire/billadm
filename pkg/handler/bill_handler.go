package handler

import (
	"fmt"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/operation"
	"ljw/billadm/utils/print"
)

var _ IResource = &BillHandler{}

type BillHandler struct {
}

func (bh *BillHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case operation.Get:
		err = bh.get(resourceName, resources, cm, options)
	case operation.Delete:
		err = bh.delete(resourceName, resources, cm, options)
	case operation.Create:
		err = bh.create(resourceName, resources, cm, options)
	case operation.Modify:
		err = bh.modify(resourceName, resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for BillHandler", op)
	}
	return err
}

func (bh *BillHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// resourceName是空的，那么我们就打印所有的bill和程序配置信息
	if resourceName == "" {
		print.PrintBillConfig(cm.Config)

		for _, bill := range cm.Config.Bills {
			print.PrintOneBill(bill, false)
		}

		return nil
	}

	// 查看resourceName存不存在，不存在就报错，否则打印详细信息
	if !cm.IsBillExist(resourceName) {
		return fmt.Errorf("bill [%s] doesn't exsit", resourceName)
	}
	bill := cm.GetBillByName(resourceName)
	print.PrintOneBill(bill, true)
	return nil
}

func (bh *BillHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (bh *BillHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (bh *BillHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
