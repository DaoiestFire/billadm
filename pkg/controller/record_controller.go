package controller

import (
	"context"
	"fmt"
	"time"

	"ljw/billadm/pkg/api/service"
	"ljw/billadm/pkg/api/v1"
)

var _ Controller = &RecordController{}

type RecordController struct {
}

func (r *RecordController) Get(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	return fmt.Errorf("not supported")
}

func (r *RecordController) Create(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	record := v1.NewRecord("")
	record.SetCost(config.Cost)
	record.SetDescription(config.Description)
	record.SetLabel(config.Label)
	record.SetConsumptionTime(time.Now().Unix())
	_, err := storageClient.CreateRecord(ctx, &service.CreateRecordRequest{Record: record.ToRecordInfo()})
	if err != nil {
		return err
	}
	return nil
}

func (r *RecordController) Delete(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	_, err := storageClient.DeleteRecord(ctx, &service.DeleteRecordRequest{DayEntryName: config.Time, Id: config.ID})
	if err != nil {
		return err
	}
	return nil
}

func (r *RecordController) Edit(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	getRecordResponse, err := storageClient.GetRecord(ctx, &service.GetRecordRequest{DayEntryName: config.Time, Id: config.ID})
	if err != nil {
		return err
	}
	record := v1.NewRecord("")
	record.FromRecordInfo(getRecordResponse.Record)
	if config.Cost > 0 {
		record.SetCost(config.Cost)
	}
	if config.Description != "" {
		record.SetDescription(config.Description)
	}
	if config.Label != 0 {
		record.SetLabel(config.Label)
	}
	_, err = storageClient.CreateRecord(ctx, &service.CreateRecordRequest{Record: record.ToRecordInfo()})
	if err != nil {
		return err
	}
	return nil
}
