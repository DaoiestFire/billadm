package manager

import (
	"encoding/json"
	"path"

	constant "ljw/billadm/const"
	"ljw/billadm/pkg/types"
	"ljw/billadm/utils/fileutils"
	"ljw/billadm/utils/pathutils"
)

var cm *ConfigManager

func Init() error {
	homePath, err := pathutils.GetHomeDir()
	if err != nil {
		return err
	}
	configPath := path.Join(homePath, constant.ConfigurationName)
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
