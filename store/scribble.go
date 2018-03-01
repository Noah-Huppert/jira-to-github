package store

import (
	"github.com/nanobox-io/golang-scribble"
)

// Scribble implements the Store interface using the Scribble library
type Scribble struct {
	// db holds the scribble database instance for the store
	db *scribble.Driver
}

// TODO Implement Store for Scribble
