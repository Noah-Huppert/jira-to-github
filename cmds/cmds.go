package cmds

import (
	"fmt"
	"github.com/urfave/cli"
)

// AllCommands returns all the commands registered in the cmds module
func AllCommands() ([]cli.Command, error) {
	// Create commands
	fetchCmd, err := NewFetchCommand()
	if err != nil {
		return nil, fmt.Errorf("error creating fetch command: %s",
			err.Error())
	}

	linkCmd, err := NewLinkCommand()
	if err != nil {
		return nil, fmt.Errorf("error creating link command: %s",
			err.Error())
	}

	createCmd, err := NewCreateCommand()
	if err != nil {
		return nil, fmt.Errorf("error creating create command: %s",
			err.Error())
	}

	// commands holds all the Commands to register with the cli library
	commands := []Command{fetchCmd, linkCmd, createCmd}

	// cmds holds all the cli library typed Command structs
	cmds := []cli.Command{}

	for _, cmd := range commands {
		cmds = append(cmds, cmd.Command())
	}

	return cmds, nil
}
