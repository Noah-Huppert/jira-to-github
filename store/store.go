package store

import (
	"encoding/json"
)

// StoreDir specifies the directory to save data in
const StoreDir string = "./data"

// Store provides an interface for managing data store information
type Store interface {
	// Set saves a value in a store under a specified id. An error is
	// returned if one occurs, nil on success.
	Set(id string, data json.Marshaler) error

	// Get retrieves a value from the store with the specified id. The data
	// along with an error is returned. This error is nil on success.
	//Get(id string, data json.Unmarshaler) error

	// Hash retrieves the hash of the key's contents
	Hash(id string) (string, error)

	// Keys returns all the keys present in a store
	Keys() ([]string, error)

	// HasKey indicates if the store contains an entry for the provided
	// key. An error is returned if one occurs.
	HasKey(id string) (bool, error)
}
