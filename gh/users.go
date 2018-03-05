package gh

import (
	"context"
	"fmt"
	"os"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/store"
)

// logger is the github logger
var logger *log.Logger = log.New(os.Stdout, "gh.users: ", 0)

// UpdateUsers retrieves all users who contribute to the GitHub repository
// specified in the configuration. And stores them in the GitHubUserStore.
//
// An error will be returned if one occurs.
func UpdateUsers(ctx context.Context, cfg *config.Config, stores *store.Stores) error {
	// GH client
	ghClient := NewClient(ctx, cfg)

	// Get repo contributors
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

	logger.Printf("contribs: %s", contribLogins)

	return nil
}
