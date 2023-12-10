package controller

import (
	"context"
	"fmt"

	"ljw/billadm/pkg/api/service"
	"ljw/billadm/pkg/api/v1"
	"ljw/billadm/utils/view"
)

var _ Controller = &DayEntryController{}

type DayEntryController struct {
}

func (d *DayEntryController) Get(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	if config.All {
		listAllDayEntryResponse, err := storageClient.ListAllDayEntry(ctx, &service.ListAllDayEntryRequest{})
		if err != nil {
			return err
		}
		dayEntryList := make([]v1.IDayEntry, 0, len(listAllDayEntryResponse.DayEntryList))
		for i := range listAllDayEntryResponse.DayEntryList {
			dayEntry := v1.NewDayEntry("")
			dayEntry.FromDayEntryInfo(listAllDayEntryResponse.DayEntryList[i])
			dayEntryList = append(dayEntryList, dayEntry)
		}
		err = view.PrintDEs(dayEntryList)
		if err != nil {
			return err
		}
		return nil
	}
	getDayEntryResponse, err := storageClient.GetDayEntry(ctx, &service.GetDayEntryRequest{Name: config.Time})
	if err != nil {
		return err
	}
	de := v1.NewDayEntry("")
	de.FromDayEntryInfo(getDayEntryResponse.DayEntry)
	records := de.ListRecords()
	err = view.PrintRecords(records)
	if err != nil {
		return err
	}
	return nil
}

func (d *DayEntryController) Create(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	createDayEntryRequest := &service.CreateDayEntryRequest{
		DayEntry: &service.DayEntryInfo{
			ObjectMeta: &service.ObjectMeta{
				Name: config.Time,
			},
		},
	}
	_, err := storageClient.CreateDayEntry(ctx, createDayEntryRequest)
	if err != nil {
		return err
	}
	return nil
}

func (d *DayEntryController) Delete(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	deleteDayEntryRequest := &service.DeleteDayEntryRequest{
		Name: config.Time,
	}
	_, err := storageClient.DeleteDayEntry(ctx, deleteDayEntryRequest)
	if err != nil {
		return err
	}
	return nil
}

func (d *DayEntryController) Edit(ctx context.Context, storageClient service.StorageServiceClient, config *v1.Config) error {
	return fmt.Errorf("not supported")
}
