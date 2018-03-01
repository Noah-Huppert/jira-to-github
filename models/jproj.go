package models

import (
	"github.com/andygrunwald/go-jira"
)

// JiraProject holds relevant Jira project information
type JiraProject struct {
	// ID holds the Jira project ID.
	ID string

	// Key holds the human readable ID of the Jira project.
	Key string

	// IssueTypes holds the names of the types of issues that can be
	// created in the project.
	IssueTypes []string
}

// NewJiraProject creates a JiraProject from a jira.Project
func NewJiraProject(from jira.Project) JiraProject {
	// Extract issue types
	types := []string{}
	for _, t := range from.IssueTypes {
		types = append(types, t.Name)
	}

	return JiraProject{
		ID:         from.ID,
		Key:        from.Key,
		IssueTypes: types,
	}
}
