package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/config"
)

func main() {
	// Logger
	logger := log.New(os.Stdout, "main: ", 0)

	// Configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// Jira
	jiraClient, err := jira.NewClient(nil, cfg.Jira.URL)
	if err != nil {
		logger.Fatalf("error creating Jira client: %s", err.Error())
	}

	jiraClient.Authentication.SetBasicAuth(cfg.Jira.Username,
		cfg.Jira.Password)

	// Get issues
	issuesQuery = fmt.Sprintf("project=%s", cfg.Jira.Project)

	issues, _, err := jiraClient.Issue.Search(issuesQuery, nil)
	if err != nil {
		logger.Fatalf("error searching for Jira issues: %s", err.Error())
	}

	for _, issue := range issues {
		logger.Print(issue)
	}
}
