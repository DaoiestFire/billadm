package v1

const (
	Bill     = "Bill"
	Record   = "Record"
	DayEntry = "DayEntry"

	V1 = "v1"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type ObjectMeta struct {
	Name string `json:"name,omitempty"`

	CreationTimestamp int64 `json:"creationTimestamp,omitempty"`
	ModifyTimestamp   int64 `json:"modifyTimestamp,omitempty"`
}
