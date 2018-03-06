package models

import (
	"encoding/json"
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/hash"
	"github.com/fatih/structs"
)

// GitHubAggregateStore indicates the store name to save GitHubAggregates into
const GitHubAggregateStore string = "gh_aggregates"

// GitHubAggregateStoreKey is the key GitHub Aggregates will always be stored under
const GitHubAggregateStoreKey string = "main"

// GitHubAggregate provides an overview of all the retrieved GitHub models
type GitHubAggregate struct {
	// Users holds the GitHubUsersAggregate
	Users *GitHubUsersAggregate

	// aggregated indicates if the data has been populated
	aggregated bool
}

// NewGitHubAggregate creates a new GitHubAggregate
func NewGitHubAggregate() *GitHubAggregate {
	// Users aggr
	users := NewGitHubUsersAggregate()

	return &GitHubAggregate{
		Users:      users,
		aggregated: false,
	}
}

func (a GitHubAggregate) String() string {
	return fmt.Sprintf("Users: {%s}\n"+
		"aggregated: %t",
		a.Users, a.aggregated)
}

// Hash implements hash.Hashable.Hash
func (a GitHubAggregate) Hash() string {
	return hash.HashStr(a.String())
}

// MarshalJSON implements json.Marshaler.MarshalJSON
func (a GitHubAggregate) MarshalJSON() ([]byte, error) {
	m := structs.Map(a)

	m["hash"] = a.Hash()

	return json.Marshal(m)
}
