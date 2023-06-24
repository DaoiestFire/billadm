package handler

import (
	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
)

var _ IResource = &LabelHandler{}

type LabelHandler struct {
}

func (lh *LabelHandler) Run(op string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}
