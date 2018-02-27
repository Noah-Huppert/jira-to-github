package cache

import (
	"fmt"
	"ioutil"
	"json"
)

// JSONStoreCacheT is the type used by JSONStore to cache data values. Keys are
// ids, values are interface{}'s.
type JSONStoreCacheT map[string]interface{}

// JSONStore implements the Store interface by saving information in JSON files.
type JSONStore struct {
	// dir indicates the directory to store the JSON file in.
	dir string

	// name is the unique store name used to identify which JSON file to
	// read in the dir path.
	name string

	// data holds the current data stored in the JSON file
	data JSONStoreCacheT

	// loaded indicates if the JSON store file has been read yet
	load bool
}

// NewJSONStore creates a JSONStore instance
func NewJSONStore(dir string, name string) *JSONStore {
	return &JSONStore{
		dir:    dir,
		name:   name,
		data:   JSONStoreCacheT{},
		loaded: false,
	}
}

// path returns the file path of the JSON file data is stored in.
func (s JSONStore) path() string {
	return fmt.Sprintf("%s/%s", s.dir, s.name)
}

// read imports the contents of the JSON store file into the data field
func (s *JSONStore) read() error {
	// Read file
	bytes, err := ioutil.ReadFile(s.path())
	if err != nil {
		return fmt.Errorf("error reading store file: %s", err.Error())
	}

	// Convert to object
	err = json.Unmarshal(bytes, &s.data)
	if err != nil {
		return fmt.Errorf("error converting store file contents "+
			"to JSON: %s", err.Error())
	}

	// Success
	return nil
}

// write exports the contents of the data field into the JSON store file
func (s JSONStore) write() error {
	// To JSON
	bytes, err := json.Marshal(s.data)
	if err != nil {
		return fmt.Errorf("error converting store data into JSON: %s",
			err.Error())
	}

	// Save
	err = ioutil.WriteFile(s.path(), bytes, 0)
	if err != nil {
		return fmt.Errorf("error saving store data: %s", err.Error())
	}

	// Success
	return nil
}

// Get implements Store.Get
func (s *JSONStore) Get(id string) (interface{}, error) {
	// If we haven't load the store file yet
	if !s.loaded {
		// Load
		if err := s.read(); err != nil {
			return fmt.Errorf("error reading store file: %s",
				err.Error())
		}
	}

	// Check if key exists
	if val, ok := self.data[id]; ok {
		// Return
		return val, nil
	} else {
		// Return nil
		return nil, nil
	}
}

// Set implements Store.Set
func (s *JSONStore) Set(id string, data interface{}) error {
	// Set key
	s.data[id] = data

	// Write
	if err := s.write(); err != nil {
		return fmt.Errorf("error writing to store file: %s",
			err.Error())
	}

	// Success
	return nil
}

// GetAll implements Store.GetAll
func (s *JSONStore) GetAll() (interface{}, error) {
	// If we haven't load the store file yet
	if !s.loaded {
		// Load
		if err := s.read(); err != nil {
			return fmt.Errorf("error reading store file: %s",
				err.Error())
		}
	}

	// Success
	return s.data, nil
}
