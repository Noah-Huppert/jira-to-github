package store

import (
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/models"
)

// GitHubAggregateStore is a store which implements an addition Get method for the
// GitHubAggregate.
type GitHubAggregateStore struct {
	*ScribbleStore
}

// NewGitHubAggregateStore creates a new GitHubAggregate store. An error is returned if one
// occurs.
func NewGitHubAggregateStore() (*GitHubAggregateStore, error) {
	// Create ScribbleStore
	str, err := NewScribbleStore(models.GitHubAggregateStore)
	if err != nil {
		return nil, fmt.Errorf("error creating base store: %s",
			err.Error())
	}

	return &GitHubAggregateStore{
		str,
	}, nil
}

// Get retrieves a GitHubAggregate
func (s GitHubAggregateStore) Get(id string, data *models.GitHubAggregate) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error retrieving GitHub aggregate: %s",
			err.Error())
	}

	return nil
}
