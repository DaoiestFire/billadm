package handler

var _ handler = &DayRecordsHandler{}

type DayRecordsHandler struct {
}

func (dh *DayRecordsHandler) Get() error {
	return nil
}

func (dh *DayRecordsHandler) Delete() error {
	return nil
}

func (dh *DayRecordsHandler) Create() error {
	return nil
}
