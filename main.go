package main

import (
	"context"
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/aggr"
	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/gh"
	"github.com/Noah-Huppert/jira-to-github/jira"
	"github.com/Noah-Huppert/jira-to-github/store"
)

func main() {
	// Logger
	logger := log.New(os.Stdout, "main: ", 0)

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

	// Load Jira issues
	if err = jira.UpdateIssues(cfg, stores); err != nil {
		logger.Fatalf("error loading Jira issues: %s", err.Error())
	}

	// Make Jira aggregate
	jAggr := aggr.NewJiraAggregate()
	if err = jAggr.Aggregate(stores); err != nil {
		logger.Fatalf("error generating Jira aggregate: %s", err.Error())
	}
	logger.Printf("Jira aggregate: %s", jAggr)

	// Load GitHub users
	if err = gh.UpdateUsers(context.Background(), cfg, stores); err != nil {
		logger.Fatalf("error loading GitHub users: %s", err.Error())
	}
}
