package v1

import (
	"encoding/json"

	constant "ljw/billadm/const"
	metav1 "ljw/billadm/pkg/api/meta/v1"
	"ljw/billadm/utils/fileutils"
	timeutils "ljw/billadm/utils/time"
)

type BillSpec struct {
	User string `json:"user,omitempty"`
}

type Bill struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BillSpec `json:"spec,omitempty"`
}

type IBill interface {
	GetName() string
	GetCreationTime() string
	GetModifyTime() string
	SetUser(user string)
	GetUser() string
}

var _ IBill = &Bill{}

func NewBill(name string) *Bill {
	bill := &Bill{}
	bill.Kind = metav1.Bill
	bill.APIVersion = metav1.V1
	bill.Name = name
	bill.CreationTimestamp = timeutils.GetNowTimeString()
	bill.ModifyTimestamp = timeutils.GetNowTimeString()
	bill.Spec.User = constant.DefaultUserName
	return bill
}

func (b *Bill) GetUser() string {
	return b.Spec.User
}

func (b *Bill) SetUser(user string) {
	b.ModifyTimestamp = timeutils.GetNowTimeString()
	b.Spec.User = user
}

func (b *Bill) GetName() string {
	return b.Name
}

func (b *Bill) GetCreationTime() string {
	return b.CreationTimestamp
}

func (b *Bill) GetModifyTime() string {
	return b.ModifyTimestamp
}

func (b *Bill) UnmarshalFrom(data []byte) error {
	return json.Unmarshal(data, b)
}

func (b *Bill) MarshalTo() ([]byte, error) {
	return fileutils.GenerateJsonData(b)
}
