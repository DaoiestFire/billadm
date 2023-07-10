package constant

import (
	"os"
)

// file
const (
	FilePerm os.FileMode = 0600
	DirPerm  os.FileMode = 0700
)

const (
	ConfigDir      = "/etc/billadm"
	ConfigName     = "billadm.config"
	BilladmDataDir = "BilladmData"

	DefaultUserName = "default"
)

// resource
const (
	Record   = "record"
	DayEntry = "de"
	Label    = "label"
	Bill     = "bill"
)

const (
	TimeFormat = "2006-01-02-15:04:05"
)

const (
	Get      = "get"
	Create   = "create"
	Delete   = "delete"
	Edit     = "edit"
	Activate = "activate"
)

const (
	LogFile         = "log_file"
	BilladmDataPath = "billadm_data_path"
)
