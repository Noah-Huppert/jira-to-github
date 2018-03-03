package store

import (
	"fmt"
)

// Stores is a collection of all the different stores used to save model data.
type Stores struct {
	// Jira holds Jira model stores
	Jira *JiraStores
}

// NewStores creates a new Stores instance. An error is returned if one occurs.
func NewStores() (*Stores, error) {
	// Jira
	jira, err := NewJiraStores()
	if err != nil {
		return nil, fmt.Errorf("error creating Jira store: %s",
			err.Error())
	}

	return &Stores{
		Jira: jira,
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
