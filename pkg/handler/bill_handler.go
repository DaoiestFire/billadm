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
		err = bh.get(resources, cm, options)
	case operation.Delete:
		err = bh.delete(resources, cm, options)
	case operation.Create:
		err = bh.create(resources, cm, options)
	case operation.Modify:
		err = bh.modify(resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for BillHandler", op)
	}
	return err
}

func (bh *BillHandler) get(resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	fmt.Println("all bills:")
	for billName, bill := range cm.Config.Bills {

	}
}

func (bh *BillHandler) delete(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (bh *BillHandler) create(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (bh *BillHandler) modify(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
