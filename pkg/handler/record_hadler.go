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

func (rh *RecordHandler) Run(op, resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case operation.Get:
		err = rh.get(resourceName, resources, cm, options)
	case operation.Delete:
		err = rh.delete(resourceName, resources, cm, options)
	case operation.Create:
		err = rh.create(resourceName, resources, cm, options)
	case operation.Modify:
		err = rh.modify(resourceName, resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for RecordHandler", op)
	}
	return err
}

func (rh *RecordHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (rh *RecordHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (rh *RecordHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (rh *RecordHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
