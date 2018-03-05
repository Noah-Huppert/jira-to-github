package cmds

import (
	"github.com/urfave/cli"
)

// Command wraps the Command() method. Which returns a cli.Command
type Command interface {
	// Command returns a cli.Command to register with the command line
	// interface router
	Command() cli.Command
}
