package jira

import (
	"fmt"
	"strconv"

	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/Noah-Huppert/jira-to-github/store"
)

// UpdateAggregate updates the Jira aggregate and stores it. An error will be
// returned if one occurs.
func UpdateAggregate(stores *store.Stores) error {
	a := models.NewJiraAggregate()

	// Issues
	issAggr, err := MakeIssuesAggregate(stores)
	if err != nil {
		return fmt.Errorf("error aggregating issues: %s", err.Error())
	}
	a.Issues = issAggr

	// Save
	if err := stores.Aggregates.Jira.Set(models.JiraAggregateStoreKey, a); err != nil {
		return fmt.Errorf("error saving aggregate to store: %s",
			err.Error())
	}

	return nil
}

// MakeIssuesAggregate returns a JiraIssuesAggregate. An error will be return
// if one occurs.
func MakeIssuesAggregate(stores *store.Stores) (*models.JiraIssuesAggregate, error) {
	a := models.NewJiraIssuesAggregate()

	// Get all Jira issue keys
	keys, err := stores.Jira.Issues.Keys()
	if err != nil {
		return nil, fmt.Errorf("error retrieving all Jira Issue keys: %s",
			err.Error())
	}

	// tempLabels is a map which temporarily records which labels the
	// issues have
	tempLabels := map[string]bool{}

	// Get each issue
	for _, key := range keys {
		var issue models.JiraIssue = models.JiraIssue{}

		err := stores.Jira.Issues.Get(key, &issue)
		if err != nil {
			return nil, fmt.Errorf("error retrieving key: %s, err: %s",
				key, err.Error())
		}

		// Add to IDs
		id, err := strconv.Atoi(issue.ID)
		if err != nil {
			return nil, fmt.Errorf("error converting issue id to int "+
				"issue: %s, err: %s", issue, err.Error())
		}
		a.IDs = append(a.IDs, id)

		// Check if greater than max
		if id > a.MaxID {
			a.MaxID = id
		}

		// Record labels
		for _, label := range issue.Labels {
			tempLabels[label] = true
		}
	}

	// Save labels
	for label := range tempLabels {
		a.Labels = append(a.Labels, label)
	}

	return a, nil
}
