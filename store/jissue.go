package store

import (
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/models"
)

// JiraIssueStore is a store which implements an addition Get method for JiraIssue
// models.
type JiraIssueStore struct {
	*ScribbleStore
}

// NewJiraIssueStore creates a new JiraIssue store. An error is returned if one occurs.
func NewJiraIssueStore() (*JiraIssueStore, error) {
	// Create ScribbleStore
	str, err := NewScribbleStore(models.JiraIssueStore)
	if err != nil {
		return nil, fmt.Errorf("error creating base store: %s",
			err.Error())
	}

	return &JiraIssueStore{
		str,
	}, nil
}

// Get retrieves a JiraIssue
func (s JiraIssueStore) Get(id string, data *models.JiraIssue) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error retrieving Jira issue: %s", err.Error())
	}

	return nil
}
