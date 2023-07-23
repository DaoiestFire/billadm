package controller

import (
	"fmt"

	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
	"ljw/billadm/utils/view"
)

var _ Controller = &LabelController{}

type LabelController struct {
}

func (l *LabelController) Get(storage *storage.Storage, config *v1.Config) error {
	err := view.PrintLabels()
	if err != nil {
		return err
	}
	return nil
}

func (l *LabelController) Create(storage *storage.Storage, config *v1.Config) error {
	return fmt.Errorf("not supported")
}

func (l *LabelController) Delete(storage *storage.Storage, config *v1.Config) error {
	return fmt.Errorf("not supported")
}

func (l *LabelController) Edit(storage *storage.Storage, config *v1.Config) error {
	return fmt.Errorf("not supported")
}
