package view

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/liushuochen/gotable"

	constant "ljw/billadm/const"
	"ljw/billadm/pkg/api/v1"
)

func PrintBills(bills []v1.IBill, currentBill string) error {
	tb, err := gotable.Create("name", "user", "creation_time", "modify_time")
	if err != nil {
		return err
	}

	for _, b := range bills {
		r := make([]string, 0, 4)
		if strings.EqualFold(b.GetName(), currentBill) {
			r = append(r, "*"+b.GetName())
		} else {
			r = append(r, b.GetName())
		}

		r = append(r, b.GetUser())
		r = append(r, time.Unix(b.GetCreationTime(), 0).Format(constant.TimeFormat))
		r = append(r, time.Unix(b.GetModifyTime(), 0).Format(constant.TimeFormat))
		err = tb.AddRow(r)
		if err != nil {
			return err
		}
	}
	fmt.Println(tb)
	return nil
}

func PrintDEs(des []v1.IDayEntry) error {
	tb, err := gotable.Create("name", "current_id", "record_nums", "creation_time", "modify_time")
	if err != nil {
		return err
	}

	for _, de := range des {
		r := make([]string, 0, 5)
		r = append(r, de.GetName())
		r = append(r, fmt.Sprintf("%03d", de.GetCurrentID()))
		r = append(r, fmt.Sprintf("%d", de.GetLen()))
		r = append(r, time.Unix(de.GetCreationTime(), 0).Format(constant.TimeFormat))
		r = append(r, time.Unix(de.GetModifyTime(), 0).Format(constant.TimeFormat))
		err = tb.AddRow(r)
		if err != nil {
			return err
		}
	}
	fmt.Println(tb)
	return nil
}

func PrintRecords(records []v1.IRecord) error {
	tb, err := gotable.Create("Id", "label", "cost", "description", "consumption_time", "creation_time", "modify_time")
	if err != nil {
		return err
	}

	for _, re := range records {
		r := make([]string, 0, 7)
		r = append(r, re.GetID())
		r = append(r, v1.LabelToChinese[re.GetLabel()])
		r = append(r, fmt.Sprintf("%f", re.GetCost()))
		r = append(r, re.GetDescription())
		r = append(r, time.Unix(re.GetConsumptionTime(), 0).Format(constant.TimeFormat))
		r = append(r, time.Unix(re.GetCreationTime(), 0).Format(constant.TimeFormat))
		r = append(r, time.Unix(re.GetModifyTime(), 0).Format(constant.TimeFormat))
		err = tb.AddRow(r)
		if err != nil {
			return err
		}
	}
	fmt.Println(tb)
	return nil
}

func PrintLabels() error {
	tb, err := gotable.Create("label_id", "label_name")
	if err != nil {
		return err
	}

	for i := 1; i < len(v1.LabelList); i++ {
		r := make([]string, 0, 2)
		r = append(r, strconv.Itoa(i))
		r = append(r, v1.LabelToChinese[v1.LabelList[i]])
		err = tb.AddRow(r)
		if err != nil {
			return err
		}
	}
	fmt.Println(tb)
	return nil
}
