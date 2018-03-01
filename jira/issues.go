package jira

import (
	"fmt"
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/andygrunwald/go-jira"
)

// Logger
var logger *log.Logger = log.New(os.Stdout, "jira: ", 0)

// UpdateIssues retrieves and saves all Jira issues that were created after
// the last sync.
//
// An error will be return if one occurs. Nil on success.
func UpdateIssues(cfg *config.Config) error {
	// Jira
	jiraClient, err := jira.NewClient(nil, cfg.Jira.URL)
	if err != nil {
		return fmt.Errorf("error creating Jira client: %s",
			err.Error())
	}

	jiraClient.Authentication.SetBasicAuth(cfg.Jira.Username,
		cfg.Jira.Password)

	// Get issues
	issuesQuery := fmt.Sprintf("project=%s", cfg.Jira.Project)
	issues, _, err := jiraClient.Issue.Search(issuesQuery, nil)
	if err != nil {
		return fmt.Errorf("error searching for Jira issues: %s",
			err.Error())
	}

	for i, issue := range issues {
		jI := models.NewJiraIssue(issue)

		logger.Printf("%d\n===\n%s\n\n", i, jI)
	}

	// Success
	return nil
}
