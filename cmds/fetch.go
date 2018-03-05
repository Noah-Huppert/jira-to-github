package cmds

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/aggr"
	"github.com/Noah-Huppert/jira-to-github/gh"
	"github.com/Noah-Huppert/jira-to-github/jira"
)

// logger is used to print messages during the fetch command execution
var logger *log.Logger = log.New(os.Stdout, "cmds.fetch: ", 0)

// FetchCommand implements Command for the fetch command. Which retrieves and
// stores state from the Jira and GitHub APIs.
type FetchCommand struct {
	*BaseCommand
}

// NewFetchCommand makes a new FetchCommand instance. An error is returned if
// one occurs.
func NewFetchCommand() (*FetchCommand, error) {
	// Base
	base, err := NewBaseCommand()
	if err != nil {
		return nil, fmt.Errorf("error creating base command: %s",
			err.Error())
	}

	return &FetchCommand{
		base,
	}, nil
}

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
	// Load Jira issues
	if err := jira.UpdateIssues(c.jiraClient, c.cfg, c.stores); err != nil {
		return fmt.Errorf("error loading Jira issues: %s", err.Error())
	}

	// Make Jira aggregate
	jAggr := aggr.NewJiraAggregate()
	if err := jAggr.Aggregate(c.stores); err != nil {
		return fmt.Errorf("error generating Jira aggregate: %s",
			err.Error())
	}
	logger.Printf("Jira aggregate: %s", jAggr)

	// Load GitHub users
	if err := gh.UpdateUsers(c.ghClient, c.ctx, c.cfg, c.stores); err != nil {
		return fmt.Errorf("error loading GitHub users: %s", err.Error())
	}

	// Make GitHub aggregate
	ghAggr := aggr.NewGitHubAggregate()
	if err := ghAggr.Aggregate(c.stores); err != nil {
		return fmt.Errorf("error generating GitHub aggregate: %s",
			err.Error())
	}
	logger.Printf("GitHub aggregate: %s", ghAggr)

	return nil
}
