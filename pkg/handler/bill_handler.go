package handler

import (
	"fmt"

	"path"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/utils/fileutils"
	"ljw/billadm/utils/print"
)

var _ IResource = &BillHandler{}

type BillHandler struct {
}

func (bh *BillHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case constant.Get:
		err = bh.get(resourceName, resources, cm, options)
	case constant.Delete:
		err = bh.delete(resourceName, resources, cm, options)
	case constant.Create:
		err = bh.create(resourceName, resources, cm, options)
	case constant.Edit:
		err = bh.modify(resourceName, resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for BillHandler", op)
	}
	return err
}

func (bh *BillHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// resourceName是空的，那么我们就打印所有的bill和程序配置信息
	if resourceName == "" {
		print.BillConfigPrint(cm.Config)

		for _, bill := range cm.Config.Bills {
			print.OneBillPrint(bill, false)
		}

		return nil
	}

	// 查看resourceName存不存在，不存在就报错，否则打印详细信息
	if !cm.IsBillExist(resourceName) {
		return fmt.Errorf("bill [%s] doesn't exsit", resourceName)
	}
	bill := cm.GetBillByName(resourceName)
	print.OneBillPrint(bill, true)
	return nil
}

func (bh *BillHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	if resourceName == "" {
		return fmt.Errorf("please specify bill name")
	}

	if !cm.IsBillExist(resourceName) {
		return fmt.Errorf("bill [%s] doesn't exist", resourceName)
	}

	billPath := path.Join(cm.Config.BillDataDir, resourceName)
	err := fileutils.RemoveDirectory(billPath)
	if err != nil {
		return err
	}

	delete(cm.Config.Bills, resourceName)

	return nil
}

// 创建一个bill。在配置中添加空的bill，创建一个bill目录
func (bh *BillHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	if cm.IsBillExist(resourceName) {
		return fmt.Errorf("can't create existed bill [%s]", resourceName)
	}

	cm.AddBill(resourceName)
	err := fileutils.CreateBillDir(resourceName)
	return err
}

// 主要是修改config中的配置
func (bh *BillHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	if resourceName == "" {
		return fmt.Errorf("please specify bill name")
	}
	if !cm.IsBillExist(resourceName) {
		return fmt.Errorf("bill [%s] doesn't exist", resourceName)
	}
	if options.User != "" {
		bill := cm.GetBillByName(resourceName)
		bill.User = options.User
	}
	return nil
}
