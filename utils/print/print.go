package print

import (
	"fmt"

	"ljw/billadm/pkg/types"
)

func BillConfigPrint(billConfig *types.BilladmConfig) {
	fmt.Printf("data saved in [%s]\n", billConfig.BillDataDir)
	fmt.Printf("config created at [%s]\n", billConfig.CreationTime)
	fmt.Printf("config modified at [%s]\n", billConfig.LastModifyTime)
	fmt.Printf("There is [%d] bill on your computer\n", len(billConfig.Bills))

	if len(billConfig.CurrentBillName) != 0 {
		fmt.Printf("******current activated bill is [%s]\n", billConfig.CurrentBillName)
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
