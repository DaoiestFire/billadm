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
	ConfigurationName = ".billadm"
	BilladmDataDir    = "BilladmData"
)

// sub command
const (
	GET    = "get"
	CREATE = "create"
	DELETE = "delete"
)

// resource
const (
	Record     = "record"
	DayRecords = "dr"
	Label      = "label"
	Bill       = "bill"
)
