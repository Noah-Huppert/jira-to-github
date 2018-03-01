package models

import (
	"github.com/andygrunwald/go-jira"
	"time"
)

// JiraIssue holds the information we can tranfer to GitHub from a Jira issue.
type JiraIssue struct {
	// ID holds the Jira issue ID.
	ID string

	// Type the style of Jira issue.
	Type string

	// Subtask indicates if the Jira Issue is a subtask
	Subtask bool

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
	Due time.Time

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
	return JiraIssue{
		ID:          from.ID,
		Type:        from.Fields.Type.Name,
		Subtask:     from.Fields.Type.Subtask,
		AssigneeKey: from.Fields.Asignee.Key,
		ProjectID:   from.Fields.Project.ID,
		Resolution:  from.Fields.Resolution.Name,
		Priority:    from.Fields.Priority.Name,
		Due:         from.Fields.Duedate,
		Status:      from.Fields.Status.Name,
		Title:       from.Fields.Summary,
		Description: from.Fields.Description,
	}
}
