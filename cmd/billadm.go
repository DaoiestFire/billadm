package main

import (
	configutils "ljw/billadm/utils/config"

	"fmt"
	"os"
	"path/filepath"

	"ljw/billadm/cmd/command"
	constant "ljw/billadm/const"
	"ljw/billadm/utils/logger"
)

func main() {
	// 尽早的初始化日志组件
	logFilePath := filepath.Join(configutils.InstallPath, constant.Log, constant.LogFile)
	logger.InitLogger(logFilePath)
	defer logger.Flush()

	cmd := command.NewBilladmCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Printf("cmd Execute failed: [%v]\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
