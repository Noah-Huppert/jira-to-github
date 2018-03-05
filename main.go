package main

import (
	"github.com/Noah-Huppert/jira-to-github/cmds"
	"github.com/urfave/cli"
	"log"
	"os"
	/*"context"
	"github.com/Noah-Huppert/jira-to-github/aggr"
	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/gh"
	"github.com/Noah-Huppert/jira-to-github/jira"
	"github.com/Noah-Huppert/jira-to-github/store"*/)

// Logger is used to print messages in the main file
var logger *log.Logger = log.New(os.Stdout, "main: ", 0)

func main() {
	// Setup cli
	app := cli.NewApp()
	app.Name = "j2gh"
	app.Usage = "Jira to GitHub issues migration tool"
	app.Version = "0.1.0"
	app.EnableBashCompletion = true
	app.Commands = cmds.AllCommands()

	// Run cmd
	if err := app.Run(os.Args); err != nil {
		logger.Fatalf("error running command: %s", err.Error())
	}

	/*ctx := context.Background()

	// Configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// Stores
	stores, err := store.NewStores()
	if err != nil {
		logger.Fatalf("error creating stores: %s", err.Error())
	}

	// Jira client
	jiraClient, err := jira.NewClient(cfg)
	if err != nil {
		logger.Fatalf("error creating Jira client: %s", err.Error())
	}

	// Load Jira issues
	if err = jira.UpdateIssues(jiraClient, cfg, stores); err != nil {
		logger.Fatalf("error loading Jira issues: %s", err.Error())
	}

	// Make Jira aggregate
	jAggr := aggr.NewJiraAggregate()
	if err = jAggr.Aggregate(stores); err != nil {
		logger.Fatalf("error generating Jira aggregate: %s",
			err.Error())
	}
	logger.Printf("Jira aggregate: %s", jAggr)

	// Load GitHub users
	ghClient := gh.NewClient(ctx, cfg)
	if err = gh.UpdateUsers(ghClient, ctx, cfg, stores); err != nil {
		logger.Fatalf("error loading GitHub users: %s", err.Error())
	}

	// Make GitHub aggregate
	ghAggr := aggr.NewGitHubAggregate()
	if err = ghAggr.Aggregate(stores); err != nil {
		logger.Fatalf("error generating GitHub aggregate: %s",
			err.Error())
	}
	logger.Printf("GitHub aggregate: %s", ghAggr)*/
}
