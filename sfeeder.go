package main

import (
	"github.com/beatoz/sfeeder/cmds"
	"os"
)

func main() {
	baseCmd := cmds.NewCmd_Base()
	if err := baseCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
