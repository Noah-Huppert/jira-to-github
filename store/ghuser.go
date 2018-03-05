package store

import (
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/models"
)

// GitHubUserStore is a store which implements an addition Get method for GitHubUser
// models.
type GitHubUserStore struct {
	*ScribbleStore
}

// NewGitHubUserStore creates a new GitHubUser store. An error is returned if one
// occurs.
func NewGitHubUserStore() (*GitHubUserStore, error) {
	// Create ScribbleStore
	str, err := NewScribbleStore(models.GitHubUserStore)
	if err != nil {
		return nil, fmt.Errorf("error creating base store: %s",
			err.Error())
	}

	return &GitHubUserStore{
		str,
	}, nil
}

// Get retrieves a GitHubUser
func (s GitHubUserStore) Get(id string, data *models.GitHubUser) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error retrieving GitHub user: %s", err.Error())
	}

	return nil
}
