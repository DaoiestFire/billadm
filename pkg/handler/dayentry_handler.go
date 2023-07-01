package handler

import (
	"fmt"

	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/operation"
	"ljw/billadm/pkg/types"
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

func (dh *DayEntryHandler) get(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (dh *DayEntryHandler) delete(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (dh *DayEntryHandler) create(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (dh *DayEntryHandler) modify(resourceName string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
