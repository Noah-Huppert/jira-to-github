package gh

import (
	"context"
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/errs"
	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/Noah-Huppert/jira-to-github/store"
	"github.com/google/go-github/github"
)

// UpdateOrgUsers retrieves all users who belongs to a specific GitHub
// organization and saves them to the GitHubUser store.
//
// An error is returned if one occurs.
func UpdateOrgUsers(ghClient *github.Client, ctx context.Context,
	cfg *config.Config, stores *store.Stores) error {

	// TODO: Implement UpdateOrgUsers
	// Get org members
	apiUsrs, resp, err := ghClient.Organizations.ListMembers(ctx,
		cfg.GitHub.RepoOwner, nil)

	// Return special error if not found
	if resp.StatusCode == 404 {
		return errs.ErrNotFound
	} else if err != nil {
		return fmt.Errorf("error retrieving organization members: %s",
			err.Error())
	}

	// Store members
	for _, apiUsr := range apiUsrs {
		// Convert to GitHub User model
		usr := models.NewGitHubUser(*apiUsr)

		// Store
		if err = stores.GitHub.Users.Set(usr.Login, usr); err != nil {
			return fmt.Errorf("error saving user to store: %s",
				err.Error())
		}
	}

	return nil
}
