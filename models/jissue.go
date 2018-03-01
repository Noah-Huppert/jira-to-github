package models

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"strings"

	"github.com/Noah-Huppert/jira-to-github/str"
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

	// Status indicates the current work state of the issue
	Status string

	// Title holds a short summary of the Jira issue
	Title string

	// Description holds a longer write up about the Jira issue
	Description string

	// Progress indicates how complete an issue is. In the range from:
	// [0, 1]
	Progress float32

	// Links holds the Jira Issue Links associated with the issue
	Links []JiraIssueLink

	// Comments holds the Jira Issue Comments associated with the issue
	Comments []JiraIssueComment

	// Labels holds the Jira labels put in the issue
	Labels []string
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

	// Parse progress
	prog := float32(from.Fields.Progress.Progress) / float32(from.Fields.Progress.Total)

	// Parse links
	links := []JiraIssueLink{}
	for _, jLink := range from.Fields.IssueLinks {
		links = append(links, NewJiraIssueLink(*jLink))
	}

	// Parse comments
	comments := []JiraIssueComment{}

	if from.Fields.Comments != nil {
		for _, jCom := range from.Fields.Comments.Comments {
			comments = append(comments, NewJiraIssueComment(*jCom))
		}
	}

	return JiraIssue{
		ID:          from.ID,
		Type:        from.Fields.Type.Name,
		AssigneeKey: assn,
		ProjectID:   from.Fields.Project.ID,
		Resolution:  res,
		Priority:    from.Fields.Priority.Name,
		Status:      from.Fields.Status.Name,
		Title:       from.Fields.Summary,
		Description: from.Fields.Description,
		Progress:    prog,
		Links:       links,
		Comments:    comments,
		Labels:      from.Fields.Labels,
	}
}

func (i JiraIssue) String() string {
	// TODO Figure out why slices not casting to slice of Stringers
	return fmt.Sprintf("%sID: %s\n"+
		"Type: %s\n"+
		"ProjectID: %s\n"+
		"AssigneeKey: %s\n"+
		"Resolution: %s\n"+
		"Priority: %s\n"+
		"Status: %s\n"+
		"Title: %s\n"+
		"Description: %s\n"+
		"Progress: %g\n"+
		"Links: [%s]\n"+
		"Comments: [%s]\n"+
		"Labels: [%s]",
		i.ID, i.Type, i.ProjectID, i.AssigneeKey, i.Resolution,
		i.Priority, i.Status, i.Title, i.Description, i.Progress,
		str.JoinStringers(i.Links, ", "), str.JoinStringers(i.Comments, ", "),
		strings.Join(i.Labels, ", "))
}
