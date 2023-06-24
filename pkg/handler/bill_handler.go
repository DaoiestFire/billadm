package handler

import (
	"ljw/billadm/cmd/options"
	"ljw/billadm/pkg/manager"
)

var _ IResource = &BillHandler{}

type BillHandler struct {
}

func (bh *BillHandler) Run(op string, resources Resources, cm *manager.ConfigManager, options *options.Options) error {
	return nil
}
