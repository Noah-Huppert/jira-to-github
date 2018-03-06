package store

import (
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/models"
)

// JiraAggregateStore is a store which implements an addition Get method for the
// JiraAggregate.
type JiraAggregateStore struct {
	*ScribbleStore
}

// NewJiraAggregateStore creates a new JiraAggregate store. An error is returned if one
// occurs.
func NewJiraAggregateStore() (*JiraAggregateStore, error) {
	// Create ScribbleStore
	str, err := NewScribbleStore(models.JiraAggregateStore)
	if err != nil {
		return nil, fmt.Errorf("error creating base store: %s",
			err.Error())
	}

	return &JiraAggregateStore{
		str,
	}, nil
}

// Get retrieves a JiraAggregate
func (s JiraAggregateStore) Get(id string, data *models.JiraAggregate) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error retrieving Jira aggregate: %s",
			err.Error())
	}

	return nil
}
