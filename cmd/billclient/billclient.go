package main

import (
	"fmt"
	"os"

	"ljw/billadm/cmd/command"
)

func main() {
	cmd := command.NewBillClientCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Printf("cmd Execute failed: [%v]\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
