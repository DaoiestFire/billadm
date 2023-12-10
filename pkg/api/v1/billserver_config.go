package v1

import (
	"github.com/spf13/pflag"
	constant "ljw/billadm/const"
	"path/filepath"
)

type BillServerConfig struct {
	InstallPath string
	SocketName  string
	LogFile     string
	SocketPath  string
	LogFilePath string
}

func NewBillServerConfig() *BillServerConfig {
	return &BillServerConfig{}
}

func (bsc *BillServerConfig) ApplyTo(fs *pflag.FlagSet) {
	fs.StringVar(&bsc.InstallPath, "install-path", "/opt/bill", "directory to install bill software")
	fs.StringVar(&bsc.SocketName, "socket-name", "billserver.sock", "specify unix socket name")
	fs.StringVar(&bsc.LogFile, "log-file", "billserver.log", "specify log file name")
}

func (bsc *BillServerConfig) Validate() error {
	bsc.SocketPath = filepath.Join(bsc.InstallPath, bsc.SocketName)
	bsc.LogFilePath = filepath.Join(bsc.InstallPath, constant.Log, bsc.LogFile)

	return nil
}
