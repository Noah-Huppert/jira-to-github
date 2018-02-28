package cache

import (
	"encoding/json"
)

// Store provides methods used to modify raw data persisted on the disk
type Store interface {
	// Get retrieves a data object by its unique id. An error is also
	// returned if one occurs, nil on success.
	Get(id string, data json.Unmarshaler) error

	// Set stores a data object with a unique id. An error is returned if
	// one occurs, nil on success.
	Set(id string, data json.Marshaler) error

	// GetAll returns all the data in the store. An error is also returned
	// if one occurs, nil on success.
	GetAll() (interface{}, error)

	// Delete removes a data object with a unique key. An error is returned
	// if one occurs, nil on success.
	Delete(id string) error
}
