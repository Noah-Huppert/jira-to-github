package store

// StoreDir specifies the directory to save data in
const StoreDir string = "./data"

// Store provides an interface for managing data store information
type Store interface {
	// Set saves a value in a store under a specified id. An error is
	// returned if one occurs, nil on success.
	Set(id string, data interface{}) error

	// Get retrieves a value from the store with the specified id. The data
	// along with an error is returned. This error is nil on success.
	Get(id string) (interface{}, error)

	// GetAll retrieves all values from a store. A data array along with an
	// error will be returned. This error will be nil on success.
	GetAll() ([]interface{}, error)
}
