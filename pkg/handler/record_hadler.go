package handler

var _ handler = &RecordHandler{}

type RecordHandler struct {
}

func (rh *RecordHandler) Get() error {
	return nil
}

func (rh *RecordHandler) Delete() error {
	return nil
}

func (rh *RecordHandler) Create() error {
	return nil
}
