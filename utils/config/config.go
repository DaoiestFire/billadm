package configutils

import (
	"fmt"
	"path"
	"sync"

	"gopkg.in/ini.v1"

	constant "ljw/billadm/const"
	"ljw/billadm/utils/fileutils"
)

var cfg *ini.File
var once sync.Once

func GetBilladmConfig() (*ini.File, error) {
	if cfg != nil {
		return cfg, nil
	}
	var err error
	once.Do(func() {
		cfgPath := path.Join(constant.ConfigDir, constant.ConfigName)
		if !fileutils.Exist(cfgPath) {
			err = fmt.Errorf("config file not found: [%s]", cfgPath)
			return
		}
		cfg, err = ini.Load(cfgPath)
	})
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
