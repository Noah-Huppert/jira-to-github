package aggr

import (
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/store"
)

// GitHubUsersAggregate aggregates all retrieved GitHub User models
type GitHubUsersAggregate struct {
	// Logins holds all the GitHubUser logins
	Logins []string

	// aggregated indicates if the data has been populated
	aggregated bool
}

// NewGitHubUsersAggregate creates a new GitHubUsersAggregate
func NewGitHubUsersAggregate() *GitHubUsersAggregate {
	return &GitHubUsersAggregate{
		Logins:     []string{},
		aggregated: false,
	}
}

func (a GitHubUsersAggregate) String() string {
	return fmt.Sprintf("Logins: %s\n"+
		"aggregated: %t",
		a.Logins, a.aggregated)
}

// Aggregate collects the GitHub information. An error is returned if one
// occurs.
func (a *GitHubUsersAggregate) Aggregate(stores *store.Stores) error {
	// Get all user keys
	keys, err := stores.GitHub.Users.Keys()
	if err != nil {
		return fmt.Errorf("error retrieving GitHub user store keys: %s",
			err.Error())
	}

	a.Logins = keys

	a.aggregated = true

	return nil
}
