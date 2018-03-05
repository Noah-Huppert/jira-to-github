package cmds

import (
	"fmt"
	"github.com/urfave/cli"
)

// CreateCommand implements Command for the create command. Which creates GitHub
// issues for the retrieved Jira issues..
type CreateCommand struct{}

// Command implements Command.Command
func (c CreateCommand) Command() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "creates GitHub issues for retrieved Jira issues",
		Action: c.Execute,
	}
}

// Execute runs when a command is invoked by the command line interface. An
// error will be returned if one occurs.
func (c CreateCommand) Execute(ctx *cli.Context) error {
	return fmt.Errorf("not implemented")
}
