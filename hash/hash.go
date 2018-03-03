package hash

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

// Hashable wraps the Hash method. Which is used to version structure field
// values.
type Hashable interface {
	// Hash returns the hash of a struct's values, along with an error if
	// one occurs.
	Hash() string
}

// HashStr hashes the provided string
func HashStr(str string) string {
	hasher := sha1.New()

	data := fmt.Sprintf("%s\n", str)
	bs := []byte(data)

	hasher.Write(bs)

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
