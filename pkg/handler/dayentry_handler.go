package handler

import (
	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/pkg/types"
)

var _ IResource = &DayEntryHandler{}

type DayEntryHandler struct {
	path     string
	dayEntry *types.DayEntry
}

func (dh *DayEntryHandler) Run(op string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}
