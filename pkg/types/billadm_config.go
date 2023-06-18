package types

type BilladmConfig struct {
	BackupDirList string `json:"backup_dir_list"`
	BillDataDir   string `json:"bill_data_dir"`

	Bills map[string]string `json:"bills"`

	CurrentBillName string `json:"current_bill_name"`

	// 2006-01-02-15:04:05
	CreationTime   string `json:"creation_time"`
	LastModifyTime string `json:"last_modify_time"`
}
