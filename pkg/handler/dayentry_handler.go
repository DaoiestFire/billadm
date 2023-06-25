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

func (dh *DayEntryHandler) Run(op string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	var err error
	switch op {
	case operation.Get:
		err = dh.get(resources, cm, options)
	case operation.Delete:
		err = dh.delete(resources, cm, options)
	case operation.Create:
		err = dh.create(resources, cm, options)
	case operation.Modify:
		err = dh.modify(resources, cm, options)
	default:
		err = fmt.Errorf("invalid op [%s] for DayEntryHandler", op)
	}
	return err
}

func (dh *DayEntryHandler) get(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (dh *DayEntryHandler) delete(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (dh *DayEntryHandler) create(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}

func (dh *DayEntryHandler) modify(resources Resources, cm *manager.ConfigManager, options *options.Options) error {

}
