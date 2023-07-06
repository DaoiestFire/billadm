package manager

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sync"

	constant "ljw/billadm/const"
	"ljw/billadm/pkg/types"
	"ljw/billadm/utils/fileutils"
	"ljw/billadm/utils/logger"
	timeutils "ljw/billadm/utils/time"
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
	cm = &ConfigManager{
		Config: &types.BilladmConfig{
			Bills: make(map[string]*types.Bill, 0),
		},
	}
	configPath := path.Join(constant.ConfigurationDir, constant.ConfigurationName)
	if !fileutils.Exist(configPath) {
		cm.Config.CreationTime = timeutils.GetNowTimeString()
		home, _ := os.UserHomeDir()
		cm.Config.BillDataDir = path.Join(home, constant.BilladmDataDir)
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

	data, err := fileutils.GenerateJsonData(cm.Config)
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
	if !cm.IsBillExist(name) {
		return fmt.Errorf("bill [%v] doesn't exist", name)
	}
	cm.Config.CurrentBillName = name
	return nil
}

func (cm *ConfigManager) IsBillExist(name string) bool {
	_, ok := cm.Config.Bills[name]
	return ok
}

func (cm *ConfigManager) GetBillByName(name string) *types.Bill {
	if !cm.IsBillExist(name) {
		return nil
	}
	return cm.Config.Bills[name]
}

func (cm *ConfigManager) AddBill(name string) {
	cm.Config.Bills[name] = types.NewBill(name, constant.DefaultUserName)
}

func (cm *ConfigManager) DeleteBill(name string) {
	delete(cm.Config.Bills, name)
}
