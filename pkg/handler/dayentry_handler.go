package handler

import "ljw/billadm/pkg/types"

type DayEntryHandler struct {
	path     string
	dayEntry *types.DayEntry
}

func (dh *DayEntryHandler) Get(am types.IArgsMap) error {
	return nil
}

func (dh *DayEntryHandler) Delete(am types.IArgsMap) error {
	return nil
}

func (dh *DayEntryHandler) Create(am types.IArgsMap) error {
	return nil
}

func (dh *DayEntryHandler) Read() error {

	return nil
}

func (dh *DayEntryHandler) Write() error {
	return nil
}
