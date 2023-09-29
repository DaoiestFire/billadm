package print

import (
	"fmt"

	v1 "ljw/billadm/pkg/api/v1"
)

func BillConfigPrint(billConfig *v1.BilladmConfig) {
	fmt.Printf("data saved in [%s]\n", billConfig.BillDataDir)
	fmt.Printf("config created at [%s]\n", billConfig.CreationTime)
	fmt.Printf("config modified at [%s]\n", billConfig.LastModifyTime)
	fmt.Printf("There is [%d] bill on your computer\n", len(billConfig.Bills))

	if len(billConfig.CurrentBillName) != 0 {
		fmt.Printf("******current activated bill is [%s]******\n", billConfig.CurrentBillName)
	} else {
		fmt.Println("!!!No activated bill!!!")
	}
}

func OneBillPrint(bill *types.Bill, detail bool) {
	if detail {
		fmt.Printf("bill: name [%s] user [%s] creation_time[%s] last_modify_time [%s]\n",
			bill.Name, bill.User, bill.CreationTime, bill.LastModifyTime)
	} else {
		fmt.Printf("bill: name [%s] user [%s] creation_time[%s] last_modify_time [%s]\n",
			bill.Name, bill.User, bill.CreationTime, bill.LastModifyTime)
	}
}

func OneRecordPrint(record types.IRecord) {
	fmt.Printf("[id -> %s] [cost -> %f] [description -> %s] [consumption time -> %s] [label -> %s]\n",
		record.GetID(),
		record.GetCost(),
		record.GetDescription(),
		record.GetTime(),
		types.LabelToChinese[record.GetLabels()])
}
