package pathutils

import (
	constant "ljw/billadm/const"
	"os"
	"os/user"
	"path"

	"ljw/billadm/pkg/manager"
)

func GetHomeDir() string {
	u, err := user.Current()
	if err != nil {
		return "/tmp"
	}
	return u.HomeDir
}

func RemoveDirectory(path string) error {
	return os.RemoveAll(path)
}

func CreateBillDir(name string) error {
	cm := manager.GetConfigManager()
	billPath := path.Join(cm.Config.BillDataDir, name)
	return createDir(billPath)
}

func createDir(path string) error {
	err := os.MkdirAll(path, constant.DirPerm)
	return err
}
