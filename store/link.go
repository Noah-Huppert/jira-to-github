package store

import (
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/models"
)

// LinkStore is a store which implements an addition Get method for Link
// models.
type LinkStore struct {
	*ScribbleStore
}

// NewLinkStore creates a new Link store. The name of the model links are being
// stored for should be provided. An error is returned if one occurs.
func NewLinkStore(model string) (*LinkStore, error) {
	// Create ScribbleStore
	str, err := NewScribbleStore(fmt.Sprintf("%s_store", model))
	if err != nil {
		return nil, fmt.Errorf("error creating base store: %s",
			err.Error())
	}

	return &LinkStore{
		str,
	}, nil
}

// Get retrieves a Link
func (s LinkStore) Get(id string, data *models.Link) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error retrieving GitHub user: %s", err.Error())
	}

	return nil
}
