package store

import (
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/models"
)

// LinkAggregateStore is a store which implements an addition Get method for LinkAggregate
// models.
type LinkAggregateStore struct {
	*ScribbleStore
}

// NewLinkAggregateStore creates a new LinkAggregate store. An error is returned if one
// occurs.
func NewLinkAggregateStore() (*LinkAggregateStore, error) {
	// Create ScribbleStore
	str, err := NewScribbleStore(models.LinkAggregateStore)
	if err != nil {
		return nil, fmt.Errorf("error creating base store: %s",
			err.Error())
	}

	return &LinkAggregateStore{
		str,
	}, nil
}

// Get retrieves a LinkAggregate
func (s LinkAggregateStore) Get(id string, data *models.LinkAggregate) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error retrieving link aggregate: %s", err.Error())
	}

	return nil
}
