package store

import (
	"fmt"
)

// Stores is a collection of all the different stores used to save model data.
type Stores struct {
	// Jira holds Jira model stores
	Jira *JiraStores

	// GitHub holds GitHub model stores
	GitHub *GitHubStores

	// Links holds Link model stores
	Links *LinkStores

	// Aggregates holds Aggregate stores
	Aggregates *AggregateStores
}

// NewStores creates a new Stores instance. An error is returned if one occurs.
func NewStores() (*Stores, error) {
	// Jira
	jira, err := NewJiraStores()
	if err != nil {
		return nil, fmt.Errorf("error creating Jira stores: %s",
			err.Error())
	}

	// GitHub
	gh, err := NewGitHubStores()
	if err != nil {
		return nil, fmt.Errorf("error creating GitHub stores: %s",
			err.Error())
	}

	// Link
	links, err := NewLinkStores()
	if err != nil {
		return nil, fmt.Errorf("error creating link model stores: %s",
			err.Error())
	}

	// Aggr
	aggrs, err := NewAggregateStores()
	if err != nil {
		return nil, fmt.Errorf("error creating aggregate stores: %s",
			err.Error())
	}

	return &Stores{
		Jira:       jira,
		GitHub:     gh,
		Links:      links,
		Aggregates: aggrs,
	}, nil
}

// JiraStores is a collection of all the stores used to save Jira model data.
type JiraStores struct {
	// Issues is the Jira Issue store
	Issues *JiraIssueStore

	// Users is the Jira User store
	Users *JiraUserStore
}

// NewJiraStores creates a new JiraStores instance. An error is returned if
// one occurs.
func NewJiraStores() (*JiraStores, error) {
	// Issues
	issues, err := NewJiraIssueStore()
	if err != nil {
		return nil, fmt.Errorf("error creating issues store: %s",
			err.Error())
	}

	// Users
	users, err := NewJiraUserStore()
	if err != nil {
		return nil, fmt.Errorf("error creating users store: %s",
			err.Error())
	}

	return &JiraStores{
		Issues: issues,
		Users:  users,
	}, nil
}

// GitHubStores is a collection of all the stores used to save GitHub model
// data.
type GitHubStores struct {
	// Users is the GitHub User store
	Users *GitHubUserStore
}

// NewGitHubStores creates a new GitHubStores instance. An error is returned
// if one occurs.
func NewGitHubStores() (*GitHubStores, error) {
	// Users
	users, err := NewGitHubUserStore()
	if err != nil {
		return nil, fmt.Errorf("error creating users store: %s",
			err.Error())
	}

	return &GitHubStores{
		Users: users,
	}, nil
}

// LinkStores is a collection of all the stores used to save Link models.
type LinkStores struct {
	// Users is the store used to save User model links
	Users *LinkStore

	// Labels is the store used to save Label model links
	Labels *LinkStore

	// Issues is the store used to save Issue model links
	Issues *LinkStore
}

// NewLinkStores create a new LinkStores instance. An error is returned if one
// occurs.
func NewLinkStores() (*LinkStores, error) {
	// Users
	users, err := NewLinkStore("users")
	if err != nil {
		return nil, fmt.Errorf("error creating a user link store: %s",
			err.Error())
	}

	// Labels
	labels, err := NewLinkStore("labels")
	if err != nil {
		return nil, fmt.Errorf("error creating a labels link store: %s",
			err.Error())
	}

	// Issues
	issues, err := NewLinkStore("issues")
	if err != nil {
		return nil, fmt.Errorf("error creating a issues link store: %s",
			err.Error())
	}

	return &LinkStores{
		Users:  users,
		Labels: labels,
		Issues: issues,
	}, nil
}

// AggregateStores is a collection of all the stores used to save aggregates.
type AggregateStores struct {
	// Jira is the store used to save Jira aggregates
	Jira *JiraAggregateStore

	// GitHub is the store used to save GitHub aggregates
	GitHub *GitHubAggregateStore

	// Links is the sotre used to save Link aggregates
	Links *LinkAggregateStore
}

// NewAggregateStores creates a new AggregateStores instance. An error is
// returned if one occurs.
func NewAggregateStores() (*AggregateStores, error) {
	// Jira
	jira, err := NewJiraAggregateStore()
	if err != nil {
		return nil, fmt.Errorf("error creating Jira aggregate store: %s",
			err.Error())
	}

	// GitHub
	gh, err := NewGitHubAggregateStore()
	if err != nil {
		return nil, fmt.Errorf("error creating GitHub aggregate store: %s",
			err.Error())
	}

	// Links
	links, err := NewLinkAggregateStore()
	if err != nil {
		return nil, fmt.Errorf("error creating Link aggregate store: %s",
			err.Error())
	}

	return &AggregateStores{
		Jira:   jira,
		GitHub: gh,
		Links:  links,
	}, nil
}
