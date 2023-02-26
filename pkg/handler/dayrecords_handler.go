package handler

import "ljw/billadm/pkg/types"

var _ handler = &DayRecordsHandler{}

type DayRecordsHandler struct {
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
