package controller

import (
	"ljw/billadm/cmd/options"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
)

var _ v1.Controller = &LabelController{}

type LabelController struct {
}

func (l *LabelController) Get(storage *storage.Storage, config *options.Config) error {

}

func (l *LabelController) Create(storage *storage.Storage, config *options.Config) error {

}

func (l *LabelController) Delete(storage *storage.Storage, config *options.Config) error {

}

func (l *LabelController) Edit(storage *storage.Storage, config *options.Config) error {

}
