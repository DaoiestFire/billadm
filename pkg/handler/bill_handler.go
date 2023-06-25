package handler

import (
	"fmt"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/operation"
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
}

func (bh *BillHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (bh *BillHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (bh *BillHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
