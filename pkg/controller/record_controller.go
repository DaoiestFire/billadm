package controller

import (
	"fmt"

	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/storage"
)

var _ Controller = &RecordController{}

type RecordController struct {
}

func (r *RecordController) Get(storage *storage.Storage, config *v1.Config) error {
	return fmt.Errorf("not supported")
}

func (r *RecordController) Create(storage *storage.Storage, config *v1.Config) error {
	record, err := storage.CreateRecord(config.Time)
	if err != nil {
		return err
	}
	record.SetCost(config.Cost)
	record.SetDescription(config.Description)
	record.SetLabel(config.Label)
	return nil
}

func (r *RecordController) Delete(storage *storage.Storage, config *v1.Config) error {
	err := storage.DeleteRecord(config.Time, config.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecordController) Edit(storage *storage.Storage, config *v1.Config) error {
	record, err := storage.GetRecord(config.Time, config.ID)
	if err != nil {
		return err
	}
	if config.Cost > 0 {
		record.SetCost(config.Cost)
	}
	if config.Description != "" {
		record.SetDescription(config.Description)
	}
	if config.Label != 0 {
		record.SetLabel(config.Label)
	}
	return nil
}
