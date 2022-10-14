package main

import (
	"os"

	"ljw/billadm/cmd/command"
	_ "ljw/billadm/utils/log"
)

func main() {
	cmd := command.NewBilladmCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
