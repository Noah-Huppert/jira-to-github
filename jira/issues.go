package jira

import (
	"fmt"
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/Noah-Huppert/jira-to-github/store"

	"github.com/andygrunwald/go-jira"
)

// logger is the Jira module logger
var logger *log.Logger = log.New(os.Stdout, "jira.issues: ", 0)

// UpdateIssues retrieves and saves all Jira issues that were created after
// the last sync.
//
// An error will be return if one occurs. Nil on success.
func UpdateIssues(jiraClient *jira.Client, cfg *config.Config,
	stores *store.Stores) error {

	// Get last Jira aggregate
	var jAggr *models.JiraAggregate = models.NewJiraAggregate()

	err := stores.Aggregates.Jira.Get(models.JiraAggregateStoreKey, jAggr)
	if err != nil {
		return fmt.Errorf("error retrieving Jira aggregate: %s", err.Error())
	}

	// Get issues
	logger.Printf("retrieving new issues (id > %d)", jAggr.Issues.MaxID)

	issuesQuery := fmt.Sprintf("project=%s", cfg.Jira.Project)
	if len(jAggr.Issues.IDs) > 0 {
		issuesQuery += fmt.Sprintf(" and id > %d", jAggr.Issues.MaxID)
	}

	issues, _, err := jiraClient.Issue.Search(issuesQuery, nil)
	if err != nil {
		return fmt.Errorf("error searching for Jira issues: %s",
			err.Error())
	}
	logger.Printf("retrieved %d new issues", len(issues))

	// Save issues
	for _, issue := range issues {
		// Create JiraIssue model from search result
		jI, err := models.NewJiraIssue(issue)
		if err != nil {
			return fmt.Errorf("error parsing Jira API response "+
				"into JiraIssue struct: %s", err.Error())
		}

		// Save JiraIssue
		if err = SaveIssue(stores, jI); err != nil {
			return fmt.Errorf("error saving Jira issue: %s",
				err.Error())
		}
	}

	// Success
	return nil
}

// SaveIssue saves a Jira Issue to a store. It also stores any new Jira Users
// which is encounters in their own Jira User store.
func SaveIssue(stores *store.Stores, issue models.JiraIssue) error {
	// Check if issue has assignee
	if issue.AssigneeKey != "" {
		// Check if assignee in store
		ok, err := stores.Jira.Users.HasKey(issue.AssigneeKey)
		if err != nil {
			return fmt.Errorf("error determining if Jira user "+
				"exists in store: %s", err.Error())
		}

		// If not stored
		if !ok {
			// Check if Jira user was parsed in issue struct
			if issue.Assignee == nil {
				// Not parsed
				return fmt.Errorf("error saving Jira issue, "+
					"encountered un-stored Jira user "+
					"with no parse information in Jira "+
					"issue, issue: %s, user key: %s",
					issue, issue.AssigneeKey)
			}

			// Store user
			if err = stores.Jira.Users.Set(issue.AssigneeKey,
				issue.Assignee); err != nil {
				return fmt.Errorf("error saving un-stored "+
					"Jira user, user key: %s, err: %s",
					issue.AssigneeKey, err.Error())
			}
			logger.Printf("stored new user, key: %s", issue.AssigneeKey)
		}
	}

	// Set issue
	if err := stores.Jira.Issues.Set(issue.ID, issue); err != nil {
		return fmt.Errorf("error saving Jira Issue: %s",
			err.Error())
	}

	return nil
}
