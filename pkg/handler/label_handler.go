package handler

import (
	"fmt"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/operation"
	"ljw/billadm/pkg/types"
)

var _ IResource = &LabelHandler{}

type LabelHandler struct {
}

func (lh *LabelHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case operation.Get:
		err = lh.get(resourceName, resources, cm, options)
	case operation.Delete:
		err = lh.delete(resourceName, resources, cm, options)
	case operation.Create:
		err = lh.create(resourceName, resources, cm, options)
	case operation.Modify:
		err = lh.modify(resourceName, resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for LabelHandler", op)
	}
	return err
}

func (lh *LabelHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 打印所有的label
	fmt.Println("here is some labels available")
	for key, label := range types.LabelToChinese {
		fmt.Printf("%s --> %b\n", label, key)
	}
	return nil
}

func (lh *LabelHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}

func (lh *LabelHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}

func (lh *LabelHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}
