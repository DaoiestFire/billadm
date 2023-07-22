package controller

import (
	"fmt"

	"ljw/billadm/cmd/options"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
	"ljw/billadm/utils/view"
)

var _ v1.Controller = &LabelController{}

type LabelController struct {
}

func (l *LabelController) Get(storage *storage.Storage, config *options.Config) error {
	err := view.PrintLabels()
	if err != nil {
		return err
	}
	return nil
}

func (l *LabelController) Create(storage *storage.Storage, config *options.Config) error {
	return fmt.Errorf("not supported")
}

func (l *LabelController) Delete(storage *storage.Storage, config *options.Config) error {
	return fmt.Errorf("not supported")
}

func (l *LabelController) Edit(storage *storage.Storage, config *options.Config) error {
	return fmt.Errorf("not supported")
}
