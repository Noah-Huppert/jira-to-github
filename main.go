package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/jira"
)

func main() {
	// Logger
	logger := log.New(os.Stdout, "main: ", 0)

	// Configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// Load Jira issues
	if err = jira.UpdateIssues(); err != nil {
		logger.Fatalf("error loading jira issues: %s", err.Error())
	}
}
