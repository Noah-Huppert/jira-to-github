package main

import (
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/config"
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
		logger.Fatalf("error loading jira issues: %s", err.Error())
	}
}
