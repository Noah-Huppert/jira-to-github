package aggr

import (
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/store"
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

// Aggregate collects the information. An error is returned if one occurs.
func (a *GitHubAggregate) Aggregate(stores *store.Stores) error {
	// Users
	if err := a.Users.Aggregate(stores); err != nil {
		return fmt.Errorf("error aggregating users: %s", err.Error())
	}

	// TODO: Save GitHubAggregate in store

	a.aggregated = true

	return nil
}
