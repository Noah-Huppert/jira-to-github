package models

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

// JiraIssue holds the information we can tranfer to GitHub from a Jira issue.
type JiraIssue struct {
	// ID holds the Jira issue ID.
	ID string

	// Type the style of Jira issue.
	Type string

	// ProjectID holds the ID of the Jira project which the issue belongs to.
	ProjectID string

	// AssigneeKey holds the Key of the Jira user who is assigned to the
	// issue
	AssigneeKey string

	// Resolution holds the final status of the Jira issue
	Resolution string

	// Priority indicate how urgent an issue is
	Priority string

	// Due indicates the date the Jira issue is due by
	Due string

	// Status indicates the current work state of the issue
	Status string

	// Title holds a short summary of the Jira issue
	Title string

	// Description holds a longer write up about the Jira issue
	Description string

	// TODO: Continue adding fields to JiraIssue struct
	// Continue after the "Status" field of https://godoc.org/github.com/andygrunwald/go-jira#IssueFields
}

// NewJiraIssue creates a new JiraIssue from a jira.Issue
func NewJiraIssue(from jira.Issue) JiraIssue {
	// Parse assignee
	assn := ""
	if from.Fields.Assignee != nil {
		assn = from.Fields.Assignee.Key
	}

	// Parse resolution
	res := ""
	if from.Fields.Resolution != nil {
		res = from.Fields.Resolution.Name
	}

	return JiraIssue{
		ID:          from.ID,
		Type:        from.Fields.Type.Name,
		AssigneeKey: assn,
		ProjectID:   from.Fields.Project.ID,
		Resolution:  res,
		Priority:    from.Fields.Priority.Name,
		Due:         from.Fields.Duedate,
		Status:      from.Fields.Status.Name,
		Title:       from.Fields.Summary,
		Description: from.Fields.Description,
	}
}

func (i JiraIssue) String() string {
	return fmt.Sprintf("ID: %s\n"+
		"Type: %s\n"+
		"ProjectID: %s\n"+
		"AssigneeKey: %s\n"+
		"Resolution: %s\n"+
		"Priority: %s\n"+
		"Due: %s\n"+
		"Status: %s\n"+
		"Title: %s\n"+
		"Description: %s",
		i.ID, i.Type, i.ProjectID, i.AssigneeKey, i.Resolution,
		i.Priority, i.Due, i.Status, i.Title, i.Description)
}
