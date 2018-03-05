package cmds

import (
	"fmt"
	"github.com/urfave/cli"
)

// FetchCommand implements Command for the fetch command. Which retrieves and
// stores state from the Jira and GitHub APIs.
type FetchCommand struct{}

// Command implements Command.Command
func (c FetchCommand) Command() cli.Command {
	return cli.Command{
		Name:   "fetch",
		Usage:  "retrieves state from the Jira and GitHub APIs",
		Action: c.Execute,
	}
}

// Execute runs when a command is invoked by the command line interface. An
// error will be returned if one occurs.
func (c FetchCommand) Execute(ctx *cli.Context) error {
	return fmt.Errorf("not implemented")
}
