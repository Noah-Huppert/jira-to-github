package gh

import (
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/store"
	"github.com/google/go-github/github"
)

// UpdateOrgUsers retrieves all users who belongs to a specific GitHub
// organization and saves them to the GitHubUser store.
//
// An error is returned if one occurs.
func UpdateOrgUsers(ghClient *github.Client, cfg *config.Config,
	stores *store.Stores) error {

	// TODO: Implement UpdateOrgUsers
	return fmt.Errorf("unimplemented")
}
