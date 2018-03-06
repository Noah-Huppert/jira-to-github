package gh

import (
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/Noah-Huppert/jira-to-github/store"
)

// UpdateAggregate updates the GitHub aggregate and stores it. An error will be
// returned if one occurs.
func UpdateAggregate(stores *store.Stores) error {
	a := models.NewGitHubAggregate()

	// Users
	users, err := MakeUserAggregate(stores)
	if err != nil {
		return fmt.Errorf("error creating user aggregate: %s",
			err.Error())
	}
	a.Users = users

	// Save
	if err = stores.Aggregates.GitHub.Set(models.GitHubAggregateStoreKey, a); err != nil {
		return fmt.Errorf("error saving aggregate to store: %s",
			err.Error())
	}

	return nil
}

// MakeUsersAggregate returns a GitHubUsersAggregate. An error will be
// returned if one occurs.
func MakeUserAggregate(stores *store.Stores) (*models.GitHubUsersAggregate, error) {
	a := models.NewGitHubUsersAggregate()

	// Get all user keys
	keys, err := stores.GitHub.Users.Keys()
	if err != nil {
		return nil, fmt.Errorf("error retrieving GitHub user store keys: %s",
			err.Error())
	}

	a.Logins = keys

	return a, nil
}
