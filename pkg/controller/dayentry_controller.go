package controller

import (
	"fmt"
	"ljw/billadm/cmd/options"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
	"ljw/billadm/utils/view"
)

var _ v1.Controller = &DayEntryController{}

type DayEntryController struct {
}

func (d *DayEntryController) Get(storage *storage.Storage, config *options.Config) error {
	if config.All {
		des, err := storage.ListAllDayEntry()
		if err != nil {
			return err
		}
		err = view.PrintDEs(des)
		if err != nil {
			return err
		}
		return nil
	}
	de, err := storage.GetDayEntry(config.Time)
	if err != nil {
		return err
	}
	records := de.ListRecords()
	err = view.PrintRecords(records)
	if err != nil {
		return err
	}
	return nil
}

func (d *DayEntryController) Create(storage *storage.Storage, config *options.Config) error {
	err := storage.CreateDayEntry(config.Time)
	if err != nil {
		return err
	}
	return nil
}

func (d *DayEntryController) Delete(storage *storage.Storage, config *options.Config) error {
	err := storage.DeleteDayEntry(config.Time)
	if err != nil {
		return err
	}
	return nil
}

func (d *DayEntryController) Edit(storage *storage.Storage, config *options.Config) error {
	return fmt.Errorf("not supported")
}
