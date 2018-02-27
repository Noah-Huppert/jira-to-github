package cache

// Store provides methods used to modify raw data persisted on the disk
type Store interface {
	// Get retrieves a data object by its unique id
	Get(id string) (interface{}, error)

	// Set stores a data object with a unique id
	Set(id string, data interface{}) error

	// GetAll returns all the data in the store
	GetAll() (interface{}, error)
}
