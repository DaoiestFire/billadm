package controller

import (
	"ljw/billadm/cmd/options"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
)

var _ v1.Controller = &DayEntryController{}

type DayEntryController struct {
}

func (d *DayEntryController) Get(storage *storage.Storage, config *options.Config) error {

}

func (d *DayEntryController) Create(storage *storage.Storage, config *options.Config) error {

}

func (d *DayEntryController) Delete(storage *storage.Storage, config *options.Config) error {

}

func (d *DayEntryController) Edit(storage *storage.Storage, config *options.Config) error {

}
