package main

import (
	"github.com/Noah-Huppert/jira-to-github/cmds"
	"github.com/urfave/cli"
	"log"
	"os"
)

// Logger is used to print messages in the main file
var logger *log.Logger = log.New(os.Stdout, "main: ", 0)

func main() {
	// Setup cli
	app := cli.NewApp()
	app.Name = "j2gh"
	app.Usage = "Jira to GitHub issues migration tool"
	app.Version = "0.1.0"
	app.EnableBashCompletion = true

	appCmds, err := cmds.AllCommands()
	if err != nil {
		logger.Fatalf("error retrieving commands: %s", err.Error())
	}
	app.Commands = appCmds

	// Run cmd
	if err := app.Run(os.Args); err != nil {
		logger.Fatalf("error running command: %s", err.Error())
	}
}
