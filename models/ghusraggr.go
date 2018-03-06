package models

import (
	"fmt"
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
