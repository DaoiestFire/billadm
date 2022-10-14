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
	LogFile string = "billadm.log"
	LogDir  string = "log"
)

const (
	DataDir      string = "data"
	MetadataFile string = "metadata.json"
)

// sub command
const (
	GET     = "get"
	CREATE  = "create"
	DELETE  = "delete"
	STATUS  = "status"
	RECOVER = "recover"
)

// resource
const (
	Label     = "label"
	Workspace = "workspace"
	Record    = "record"
)
