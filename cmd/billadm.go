package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	constant "ljw/billadm/const"
	configutils "ljw/billadm/utils/config"
	"ljw/billadm/utils/logger"
	"os"
	"path"

	"ljw/billadm/cmd/command"
)

func main() {
	cfg, err := configutils.GetBilladmConfig()
	if err != nil {
		fmt.Printf("GetBilladmConfig failed --> <%v>\n", err)
		os.Exit(1)
	}

	// 尽早的初始化日志组件
	logFile := cfg.Section(ini.DefaultSection).Key(constant.LogFile).String()
	billadmDatePath := cfg.Section(ini.DefaultSection).Key(constant.BilladmDataPath).String()
	logFilePath := path.Join(billadmDatePath, constant.BilladmDataDir, logFile)
	logger.InitLogger(logFilePath)
	defer logger.Flush()

	cmd := command.NewBilladmCommand()
	if err := cmd.Execute(); err != nil {

	}
	os.Exit(0)
}
