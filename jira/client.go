package jira

import (
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/config"
	"github.com/andygrunwald/go-jira"
)

// NewClient creates a new Jira Client. An error is returned if one occurs.
func NewClient(cfg *config.Config) (*jira.Client, error) {
	jiraClient, err := jira.NewClient(nil, cfg.Jira.URL)
	if err != nil {
		return nil, fmt.Errorf("error creating Jira client: %s",
			err.Error())
	}

	jiraClient.Authentication.SetBasicAuth(cfg.Jira.Username,
		cfg.Jira.Password)

	return jiraClient, nil
}
