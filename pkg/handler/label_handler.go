package handler

var _ handler = &LabelHandler{}

type LabelHandler struct {
}

func (lh *LabelHandler) Get() error {
	return nil
}

func (lh *LabelHandler) Delete() error {
	return nil
}

func (lh *LabelHandler) Create() error {
	return nil
}
