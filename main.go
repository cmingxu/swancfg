package main

import (
	"fmt"
	"github.com/Dataman-Cloud/swancfg/command"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "swancfg"
	app.Usage = "command-line client for swan"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		command.NewQuotaCommand(),
		command.NewAddCommand(),
		command.NewRunCommand(),
		command.NewListCommand(),
		command.NewDeleteCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
