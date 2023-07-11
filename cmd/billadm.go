package main

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/ini.v1"

	"ljw/billadm/cmd/command"
	constant "ljw/billadm/const"
	configutils "ljw/billadm/utils/config"
	"ljw/billadm/utils/logger"
)

func main() {
	cfg, err := configutils.GetBilladmConfig()
	if err != nil {
		fmt.Printf("GetBilladmConfig failed --> <%v>\n", err)
		os.Exit(1)
	}

	// 尽早的初始化日志组件
	logFile := cfg.Section(ini.DefaultSection).Key(constant.LogFileKey).String()
	billadmDatePath := cfg.Section(ini.DefaultSection).Key(constant.BilladmDataPathKey).String()
	logFilePath := path.Join(billadmDatePath, constant.BilladmData, logFile)
	logger.InitLogger(logFilePath)
	defer logger.Flush()

	cmd := command.NewBilladmCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Printf("cmd Execute failed --> <%v>\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
