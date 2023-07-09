package handler

import (
	"fmt"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/types"
)

var _ IResource = &LabelHandler{}

type LabelHandler struct {
}

func (lh *LabelHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case constant.Get:
		err = lh.get(resourceName, resources, cm, options)
	case constant.Delete:
		err = lh.delete(resourceName, resources, cm, options)
	case constant.Create:
		err = lh.create(resourceName, resources, cm, options)
	case constant.Edit:
		err = lh.edit(resourceName, resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for LabelHandler", op)
	}
	return err
}

func (lh *LabelHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	// 打印所有的label
	fmt.Println("here is some labels available")
	for idx, label := range types.LabelList {
		if idx == 0 {
			continue
		}
		fmt.Printf("%d --> %s\n", idx, types.LabelToChinese[label])
	}
	return nil
}

func (lh *LabelHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}

func (lh *LabelHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}

func (lh *LabelHandler) edit(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}
