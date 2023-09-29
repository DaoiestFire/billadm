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
	Bin  = "bin"
	Data = "data"
	Log  = "log"

	LogFile         = "billadm.log"
	BillConfig      = "bill.json"
	CurrentBillName = "current_bill_name"
)

const (
	DefaultUserName = "default"
)

// resource
const (
	Record   = "record"
	DayEntry = "de"
	Label    = "label"
	Bill     = "bill"
)

// sub command
const (
	Get      = "get"
	Create   = "create"
	Delete   = "delete"
	Edit     = "edit"
	Activate = "activate"
)

const (
	TimeFormat = "2006-01-02-15:04:05"
)
