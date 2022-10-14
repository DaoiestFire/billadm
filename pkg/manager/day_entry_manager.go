package manager

import "ljw/billadm/pkg/types"

var _ manager = &DayEntryManager{}

type DayEntryManager struct {
	data types.Data
}

func (dem *DayEntryManager) Get() error {
	return nil
}

func (dem *DayEntryManager) Delete() error { 
	return nil
}

func (dem *DayEntryManager) Create() error {
	return nil
}

func (dem *DayEntryManager) Status() error {
	return nil
}

func (dem *DayEntryManager) Recover() error {
	return nil
}
