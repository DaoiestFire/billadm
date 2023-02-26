package handler

type handler interface {
	Get() error
	Delete() error
	Create() error
}
