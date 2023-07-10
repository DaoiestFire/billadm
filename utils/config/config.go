package configutils

import (
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
		cfg, err = ini.Load(cfgPath)
	})
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
