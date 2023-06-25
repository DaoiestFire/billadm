package handler

import (
	"fmt"
	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/operation"
)

var _ IResource = &LabelHandler{}

type LabelHandler struct {
}

func (lh *LabelHandler) Run(op string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case operation.Get:
		err = lh.get(resources, cm, options)
	case operation.Delete:
		err = lh.delete(resources, cm, options)
	case operation.Create:
		err = lh.create(resources, cm, options)
	case operation.Modify:
		err = lh.modify(resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for LabelHandler", op)
	}
	return err
}

func (lh *LabelHandler) get(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (lh *LabelHandler) delete(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (lh *LabelHandler) create(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (lh *LabelHandler) modify(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
