package handler

import (
	"fmt"
	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/operation"
)

var _ IResource = &RecordHandler{}

type RecordHandler struct {
}

func (rh *RecordHandler) Run(op string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case operation.Get:
		err = rh.get(resources, cm, options)
	case operation.Delete:
		err = rh.delete(resources, cm, options)
	case operation.Create:
		err = rh.create(resources, cm, options)
	case operation.Modify:
		err = rh.modify(resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for RecordHandler", op)
	}
	return err
}

func (rh *RecordHandler) get(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (rh *RecordHandler) delete(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (rh *RecordHandler) create(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (rh *RecordHandler) modify(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
