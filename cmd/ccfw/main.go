package main

import (
	"os"

	"github.com/2754github/ccfw/cmd/ccfw/subcmd"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		subcmd.Init()
		subcmd.Sync()

		return
	}

	switch args[0] {
	case "init":
		subcmd.Init()
	case "sync":
		subcmd.Sync()
	default:
		subcmd.Help()
	}
}
