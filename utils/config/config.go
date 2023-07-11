package configutils

import (
	"ljw/billadm/utils/fileutils"
	"os"
	"path"
	"sync"

	"gopkg.in/ini.v1"

	constant "ljw/billadm/const"
)

var cfg *ini.File
var once *sync.Once

func GetBilladmConfig() (*ini.File, error) {
	if cfg != nil {
		return cfg, nil
	}
	var err error
	once.Do(func() {
		cfgPath := path.Join(constant.ConfigDir, constant.ConfigName)
		if !fileutils.Exist(cfgPath) {
			err = fileutils.CreateDirectory(cfgPath)
			if err != nil {
				return
			}
			homePath, errTmp := os.UserHomeDir()
			if errTmp != nil {
				err = errTmp
				return
			}
			cfg = ini.Empty()
			_, err = cfg.Section(ini.DefaultSection).NewKey(constant.LogFileKey, constant.LogFile)
			if err != nil {
				return
			}
			_, err = cfg.Section(ini.DefaultSection).NewKey(constant.BilladmDataPathKey, homePath)
			if err != nil {
				return
			}
			err = cfg.SaveTo(cfgPath)
			return
		}
		cfg, err = ini.Load(cfgPath)
	})
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
