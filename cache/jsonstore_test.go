package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// testStoreT is the type for objects stored in test_data/teststore.json
type testStoreObj struct {
	a int
	b int
	c bool
}

// newTestStoreObj creates a new testStoreObj from a map[string]interface{}
func newTestStoreObj(d map[string]interface)

// TestGetExists ensures JSONStore.Get returns a object when it exists in the
// store
func TestGetExists(t *testing.T) {
	s := NewJSONStore("test_data", "teststore")

	// Get value
	obj, err := s.Get("exists")
	if err != nil {
		t.Fatalf("error calling Store.Get: %s", err.Error())
	}

	actual := obj.(testStoreObj)

	// Test
	expected := testStoreObj{
		a: 1,
		b: 2,
		c: true,
	}
	assert.EqualValues(t, expected, actual)
}
