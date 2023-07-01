package types

type Bill struct {
	Name string `json:"name"`
	User string `json:"user"`

	// 2006-01-02-15:04:05
	CreationTime   string `json:"creation_time"`
	LastModifyTime string `json:"last_modify_time"`
	LastBackupTime string `json:"last_backup_time,omitempty"`

	// how many days
	BackupTimeInterval int `json:"backup_time_interval,omitempty"`
}
