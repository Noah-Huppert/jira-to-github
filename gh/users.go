package gh

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/errs"
	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/Noah-Huppert/jira-to-github/store"
	"github.com/google/go-github/github"
)

// logger is the github logger
var logger *log.Logger = log.New(os.Stdout, "gh.users: ", 0)

// UpdateUsers retrieves all users who contribute to the GitHub repository
// specified in the configuration. And stores them in the GitHubUserStore.
//
// An error will be returned if one occurs.
func UpdateUsers(ghClient *github.Client, ctx context.Context,
	cfg *config.Config, stores *store.Stores) error {

	// Get repo contributor logins
	contribs, _, err := ghClient.Repositories.ListContributors(
		ctx, cfg.GitHub.RepoOwner, cfg.GitHub.RepoName, nil)
	if err != nil {
		return fmt.Errorf("error retrieving repository "+
			"contributors: %s", err.Error())
	}

	contribLogins := []string{}
	for _, contrib := range contribs {
		contribLogins = append(contribLogins, *contrib.Login)
	}

	// Get GitHub user for each contributor login
	for _, login := range contribLogins {
		// Get user from GitHub API
		apiUsr, _, err := ghClient.Users.Get(ctx, login)
		if err != nil {
			return fmt.Errorf("error retrieving user, login: %s"+
				", err: %s", err.Error())
		}

		// Convert to GitHubUser model
		usr := models.NewGitHubUser(*apiUsr)

		// Save in GitHub user store
		if err = stores.GitHub.Users.Set(usr.Login, usr); err != nil {
			return fmt.Errorf("error storing user, user: %s, "+
				"err: %s", usr, err.Error())
		}
	}

	// Get contributors for organization if cfg.GitHub.RepoOwner is an
	// organization
	err = UpdateOrgUsers(ghClient, ctx, cfg, stores)

	// Ignore not found error, means cfg.GitHub.RepoOwner is not an org
	if (err != nil) && (err != errs.ErrNotFound) {
		return fmt.Errorf("error updating organization users: %s",
			err.Error())
	}

	return nil
}
