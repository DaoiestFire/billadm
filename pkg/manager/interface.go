package manager

type manager interface {
	Get() error
	Delete() error
	Create() error
	Status() error
	Recover() error
}
