package handler

import (
	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
)

var _ IResource = &RecordHandler{}

type RecordHandler struct {
}

func (rh *RecordHandler) Run(op string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}
