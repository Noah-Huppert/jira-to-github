package models

import (
	"github.com/andygrunwald/go-jira"
)

// JiraUser holds information about a Jira user
type JiraUser struct {
	// Key is a unique user handle
	Key string

	// Name is the user's name
	Name string

	// Email is the user's email address
	Email string

	// DisplayName is the name the user chooses to display
	DisplayName string
}

// NewJiraUser creates a new JiraUser from a jira.User
func NewJiraUser(from jira.User) JiraUser {
	return JiraUser{
		Key:         from.Key,
		Name:        from.Name,
		Email:       from.EmailAddress,
		DisplayName: from.DisplayName,
	}
}
