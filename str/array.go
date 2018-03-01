package str

import (
	"fmt"
	"strings"
)

// JoinStringers joins a slice of Stringer interfaces with a delimiter.
func JoinStringers(items []fmt.Stringer) string {
	strs := []string{}
	for _, item := range items {
		strs = append(strs, item.String())
	}

	return fmt.Sprintf("[{%s}]",
		strings.Join(strs, "}, {"))
}
