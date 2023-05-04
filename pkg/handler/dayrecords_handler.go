package handler

import "ljw/billadm/pkg/types"

var _ handler = &DayRecordsHandler{}

type DayRecordsHandler struct {
	path     string
	dayEntry *types.DayEntry
}

func (dh *DayRecordsHandler) Get(am types.IArgsMap) error {
	return nil
}

func (dh *DayRecordsHandler) Delete(am types.IArgsMap) error {
	return nil
}

func (dh *DayRecordsHandler) Create(am types.IArgsMap) error {
	return nil
}

func (dh *DayRecordsHandler) Read() error {

	return nil
}

func (dh *DayRecordsHandler) Write() error {
	return nil
}
