package constant

import (
	"os"
)

// file variables
const (
	FilePerm os.FileMode = 0600
	DirPerm  os.FileMode = 0700
)

const (
	ConfigurationDir  = "/etc/billadm"
	ConfigurationName = "billadm.config"
	BilladmDataDir    = "BilladmData"
)

// resource
const (
	Record   = "record"
	DayEntry = "de"
	Label    = "label"
	Bill     = "bill"
)

// log
const (
	LogFormat  = "%s [%s] {%s}"
	TimeFormat = "2006-01-02-15:04:05"
)
