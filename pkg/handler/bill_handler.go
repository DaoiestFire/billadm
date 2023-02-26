package handler

var _ handler = &BillHandler{}

type BillHandler struct {
}

func (bh *BillHandler) Get() error {
	return nil
}

func (bh *BillHandler) Delete() error {
	return nil
}

func (bh *BillHandler) Create() error {
	return nil
}
