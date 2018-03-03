package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nanobox-io/golang-scribble"
	"io/ioutil"
	"os"
	"strings"
)

// ScribbleStore implements the Store interface using the Scribble library
type ScribbleStore struct {
	// db holds the scribble database instance for the store
	db *scribble.Driver

	// name is the table name of the store
	name string
}

// NewScribbleStore creates a new ScribbleStore for the specified store name.
// Returns an error if one occurs, nil on success.
func NewScribbleStore(name string) (*ScribbleStore, error) {
	// Make scribble db
	db, err := scribble.New(StoreDir, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating scribble db: %s",
			err.Error())
	}

	return &ScribbleStore{
		db:   db,
		name: name,
	}, nil
}

// Set implements Store.Set
func (s ScribbleStore) Set(id string, data json.Marshaler) error {
	if err := s.db.Write(s.name, id, data); err != nil {
		return fmt.Errorf("error saving data: %s", err.Error())
	}

	return nil
}

// Get implements Store.Get
/*func (s ScribbleStore) Get(id string, data json.Unmarshaler) error {
	if err := s.db.Read(s.name, id, data); err != nil {
		return fmt.Errorf("error reading data: %s", err.Error())
	}

	return nil
}*/

// Hash implements Store.Hash
func (s ScribbleStore) Hash(id string) (string, error) {
	data := map[string]interface{}{}

	// Get
	if err := s.db.Read(s.name, id, &data); err != nil {
		return "", fmt.Errorf("error reading data: %s", err.Error())
	}

	// Check has hash
	if hash, ok := data["hash"]; ok {
		if str, ok := hash.(string); ok {
			return str, nil
		} else {
			return "", errors.New("cannot convert hash into string")
		}
	} else {
		return "", errors.New("key has no hash stored")
	}
}

// Keys implements Store.Keys
func (s ScribbleStore) Keys() ([]string, error) {
	dir := fmt.Sprintf("%s/%s", StoreDir, s.name)

	// Check dir exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Doesn't exist, Empty
		return []string{}, nil
	}

	// Get files
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading store directory: %s", err.Error())
	}

	names := []string{}

	for _, file := range files {
		nameParts := strings.Split(file.Name(), ".")
		name := strings.Join(nameParts[0:len(nameParts)-1], ".")
		names = append(names, name)
	}

	return names, nil
}

// HasKey implements Store.HasKey
func (s ScribbleStore) HasKey(id string) (bool, error) {
	// Get keys
	keys, err := s.Keys()
	if err != nil {
		return false, fmt.Errorf("error retrieving Jira issue store "+
			"keys: %s", err.Error())
	}

	// Search
	for _, key := range keys {
		if key == id {
			// Found
			return true, nil
		}
	}

	// Not found
	return false, nil
}
