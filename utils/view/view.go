package view

import (
	"fmt"

	"github.com/liushuochen/gotable"

	v1 "ljw/billadm/pkg/api/v1"
)

func PrintBills(bills []*v1.Bill) error {
	tb, err := gotable.Create("name", "user", "creation_time", "modify_time")
	if err != nil {
		return err
	}

	for _, b := range bills {
		r := make([]string, 0, 4)
		r = append(r, b.Name)
		r = append(r, b.Spec.User)
		r = append(r, b.CreationTimestamp)
		r = append(r, b.ModifyTimestamp)
		err = tb.AddRow(r)
		if err != nil {
			return err
		}
	}
	fmt.Println(tb)
	return nil
}

func PrintDEs(des []v1.IDayEntry) error {
	tb, err := gotable.Create("name", "curren_id", "record_nums", "creation_time", "modify_time")
	if err != nil {
		return err
	}

	for _, de := range des {
		r := make([]string, 0, 5)
		r = append(r, de.GetName())
		r = append(r, fmt.Sprintf("%03d", de.GetCurrentID()))
		r = append(r, fmt.Sprintf("%d", de.GetLen()))
		r = append(r, de.GetCreationTime())
		r = append(r, de.GetModifyTime())
		err = tb.AddRow(r)
		if err != nil {
			return err
		}
	}
	fmt.Println(tb)
	return nil
}

func PrintRecords(records []v1.IRecord) error {
	tb, err := gotable.Create("ID", "label", "cost", "description", "consumption_time", "creation_time", "modify_time")
	if err != nil {
		return err
	}

	for _, re := range records {
		r := make([]string, 0, 7)
		r = append(r, re.GetID())
		r = append(r, v1.LabelToChinese[re.GetLabel()])
		r = append(r, fmt.Sprintf("%f", re.GetCost()))
		r = append(r, re.GetDescription())
		r = append(r, re.GetConsumptionTime())
		r = append(r, re.GetCreationTime())
		r = append(r, re.GetModifyTime())
		err = tb.AddRow(r)
		if err != nil {
			return err
		}
	}
	fmt.Println(tb)
	return nil
}
