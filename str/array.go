package str

import (
	"fmt"
	"strings"
)

// JoinStringers joins a slice of Stringer interfaces with a delimiter.
func JoinStringers(items []fmt.Stringer, delim string) string {
	strs := []string{}
	for _, item := range items {
		strs = append(strs, item.String())
	}

	return strings.Join(strs, delim)
}
