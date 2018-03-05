package cmds

import (
	"context"
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/Noah-Huppert/jira-to-github/gh"
	"github.com/Noah-Huppert/jira-to-github/jira"
	"github.com/Noah-Huppert/jira-to-github/store"

	jiraLib "github.com/andygrunwald/go-jira"
	"github.com/google/go-github/github"
)

// BaseCommand provides several basic fields which are generally always needed
// in command implementations.
type BaseCommand struct {
	// ctx holds the Context used to control sub-process execution
	ctx context.Context

	// cfg holds application configuration
	cfg *config.Config

	// stores holds all the stores used to save all the models
	stores *store.Stores

	// jiraClient is the Jira API client used to make requests
	jiraClient *jiraLib.Client

	// ghClient is the GitHub API client used to make requests
	ghClient *github.Client
}

// NewBaseCommand creates a new BaseCommand instance
func NewBaseCommand() (*BaseCommand, error) {
	// Context
	ctx := context.Background()

	// Configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading configuration: %s",
			err.Error())
	}

	// Stores
	stores, err := store.NewStores()
	if err != nil {
		return nil, fmt.Errorf("error creating stores: %s", err.Error())
	}

	// Jira client
	jiraClient, err := jira.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating Jira client: %s", err.Error())
	}

	// GitHub client
	ghClient := gh.NewClient(ctx, cfg)

	return &BaseCommand{
		ctx:        ctx,
		cfg:        cfg,
		stores:     stores,
		jiraClient: jiraClient,
		ghClient:   ghClient,
	}, nil
}
