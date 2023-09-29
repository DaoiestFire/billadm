package v1

type Config struct {
	Name string
	User string
	Time string
	All  bool

	Cost        float32
	Description string
	Label       LabelType
	ID          string
}
