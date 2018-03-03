package store

import (
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/models"
)

// JiraUserStore is a store which implements an addition Get method for JiraUser
// models.
type JiraUserStore struct {
	*ScribbleStore
}

// NewJiraUserStore creates a new JiraUser store. An error is returned if one
// occurs.
func NewJiraUserStore() (*JiraUserStore, error) {
	// Create ScribbleStore
	str, err := NewScribbleStore(models.JiraUserStore)
	if err != nil {
		return nil, fmt.Errorf("error creating base store: %s",
			err.Error())
	}

	return &JiraUserStore{
		str,
	}, nil
}

// Get saves a JiraUser to a store
func (s JiraUserStore) Get(id string, data *models.JiraUser) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error retrieving Jira user: %s", err.Error())
	}

	return nil
}
