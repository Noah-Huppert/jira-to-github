package cmds

import (
	"github.com/urfave/cli"
)

// commands holds all the Commands to provide via the cli
var commands []Command = []Command{
	FetchCommand{}, LinkCommand{}, CreateCommand{},
}

// AllCommands returns all the commands registered in the cmds module
func AllCommands() []cli.Command {
	cmds := []cli.Command{}

	for _, cmd := range commands {
		cmds = append(cmds, cmd.Command())
	}

	return cmds
}
