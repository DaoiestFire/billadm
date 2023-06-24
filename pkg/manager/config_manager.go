package manager

import (
	"encoding/json"
	"fmt"
	"path"
	"sync"

	constant "ljw/billadm/const"
	"ljw/billadm/pkg/types"
	"ljw/billadm/utils/fileutils"
	"ljw/billadm/utils/logger"
)

var cm *ConfigManager
var once sync.Once

func GetConfigManager() *ConfigManager {
	once.Do(func() {
		err := Init()
		if err != nil {
			logger.Errorf("init ConfigManager failed : %v", err)
		}
	})
	return cm
}

func Init() error {
	configPath := path.Join(constant.ConfigurationDir, constant.ConfigurationName)
	data, err := fileutils.ReadFileByte(configPath)
	if err != nil {
		return err
	}
	cm = &ConfigManager{}
	if err = json.Unmarshal(data, cm.Config); err != nil {
		return nil
	}

	return nil
}

type ConfigManager struct {
	Config *types.BilladmConfig
}

func (cm *ConfigManager) Save() error {
	configPath := path.Join(constant.ConfigurationDir, constant.ConfigurationName)

	data, err := json.MarshalIndent(cm.Config, "  ", "  ")
	if err != nil {
		return fmt.Errorf("marshal billadm.config failed -> <%v>", err)
	}
	if err = fileutils.WriteFileByte(configPath, data); err != nil {
		return err
	}
	return nil
}
