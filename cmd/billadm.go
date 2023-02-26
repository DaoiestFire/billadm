package main

import (
	"os"

	"ljw/billadm/cmd/command"
)

func main() {
	cmd := command.NewBilladmCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
