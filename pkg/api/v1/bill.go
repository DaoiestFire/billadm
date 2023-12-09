package v1

import (
	"encoding/json"
	"time"

	constant "ljw/billadm/const"
	metav1 "ljw/billadm/pkg/api/meta/v1"
	"ljw/billadm/pkg/api/service"
	"ljw/billadm/utils/fileutils"
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
	GetCreationTime() int64
	GetModifyTime() int64
	GetUser() string
	SetUser(string)

	UnmarshalFrom([]byte) error
	MarshalTo() ([]byte, error)

	ToBillInfo() *service.BillInfo

	Clone() IBill
}

var _ IBill = &Bill{}

func NewBill(name string) IBill {
	bill := &Bill{}
	bill.Kind = metav1.Bill
	bill.APIVersion = metav1.V1
	bill.Name = name
	bill.CreationTimestamp = time.Now().Unix()
	bill.ModifyTimestamp = time.Now().Unix()
	bill.Spec.User = constant.DefaultUserName
	return bill
}

func (b *Bill) GetName() string {
	return b.Name
}

func (b *Bill) GetCreationTime() int64 {
	return b.CreationTimestamp
}

func (b *Bill) GetModifyTime() int64 {
	return b.ModifyTimestamp
}

func (b *Bill) GetUser() string {
	return b.Spec.User
}

func (b *Bill) SetUser(user string) {
	b.ModifyTimestamp = time.Now().Unix()
	b.Spec.User = user
}

func (b *Bill) ToBillInfo() *service.BillInfo {
	return &service.BillInfo{
		TypeMeta: &service.TypeMeta{
			Kind:       b.Kind,
			ApiVersion: b.APIVersion,
		},
		ObjectMeta: &service.ObjectMeta{
			Name:              b.Name,
			CreationTimestamp: b.CreationTimestamp,
			ModifyTimestamp:   b.ModifyTimestamp,
		},
		User: b.Spec.User,
	}
}

func (b *Bill) Clone() IBill {
	return &Bill{
		TypeMeta: metav1.TypeMeta{
			Kind:       b.Kind,
			APIVersion: b.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              b.Name,
			CreationTimestamp: b.CreationTimestamp,
			ModifyTimestamp:   b.ModifyTimestamp,
		},
		Spec: BillSpec{
			User: b.Spec.User,
		},
	}
}

func (b *Bill) UnmarshalFrom(data []byte) error {
	return json.Unmarshal(data, b)
}

func (b *Bill) MarshalTo() ([]byte, error) {
	return fileutils.GenerateJsonData(b)
}
