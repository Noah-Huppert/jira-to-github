package cmds

import (
	"fmt"
	"github.com/urfave/cli"
)

// LinkCommand implements Command for the link command. Which creates a
// relationship between a Jira and GitHub API entity.
type LinkCommand struct {
	*BaseCommand
}

// NewLinkCommand creates a new LinkCommand instance. An error is returned
// if one occurs.
func NewLinkCommand() (*LinkCommand, error) {
	// Base
	base, err := NewBaseCommand()
	if err != nil {
		return nil, fmt.Errorf("error creating base command: %s",
			err.Error())
	}

	return &LinkCommand{
		base,
	}, nil
}

// Command implements Command.Command
func (c LinkCommand) Command() cli.Command {
	return cli.Command{
		Name:   "link",
		Usage:  "creates an association between a Jira and GitHub entity",
		Action: c.Execute,
	}
}

// Execute runs when a command is invoked by the command line interface. An
// error will be returned if one occurs.
func (c LinkCommand) Execute(ctx *cli.Context) error {
	return fmt.Errorf("not implemented")
}
