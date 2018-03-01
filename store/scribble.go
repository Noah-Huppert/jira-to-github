package store

import (
	"encoding/json"
	"fmt"
	"github.com/nanobox-io/golang-scribble"
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
func (s ScribbleStore) Set(id string, data interface{}) error {
	if err := s.db.Write(s.name, id, data); err != nil {
		return fmt.Errorf("error saving data: %s", err.Error())
	}

	return nil
}

// Get implements Store.Get
func (s ScribbleStore) Get(id string) (interface{}, error) {
	var data interface{}
	if err := s.db.Read(s.name, id, &data); err != nil {
		return nil, fmt.Errorf("error reading data: %s", err.Error())
	}

	return data, nil
}

// GetAll implements Store.GetAll
func (s ScribbleStore) GetAll() ([]interface{}, error) {
	rows, err := s.db.ReadAll(s.name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving all data objects: %s",
			err.Error())
	}

	items := []interface{}{}
	for _, row := range rows {
		var item interface{}

		err = json.Unmarshal([]byte(row), &item)
		if err != nil {
			return nil, fmt.Errorf("error converting row into "+
				"object, row: %s, err: %s", row, err.Error())
		}

		items = append(items, item)
	}

	return items, nil
}
