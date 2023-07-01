package manager

import (
	"encoding/json"
	"fmt"
	timeutils "ljw/billadm/utils/time"
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
	cm = &ConfigManager{}
	configPath := path.Join(constant.ConfigurationDir, constant.ConfigurationName)
	if !fileutils.Exist(configPath) {
		cm.Config.CreationTime = timeutils.GetNowTimeString()
		return nil
	}
	data, err := fileutils.ReadFileByte(configPath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, cm.Config); err != nil {
		return nil
	}

	return nil
}

type ConfigManager struct {
	Config *types.BilladmConfig
}

func (cm *ConfigManager) Save() error {
	cm.Config.LastModifyTime = timeutils.GetNowTimeString()
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

func (cm *ConfigManager) SetCurrentBillName(name string) error {
	// 需要判断name存不存在
	if _, ok := cm.Config.Bills[name]; !ok {
		return fmt.Errorf("bill [%v] doesn't exist", name)
	}
	cm.Config.CurrentBillName = name
	return nil
}
