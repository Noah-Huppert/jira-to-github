package models

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

// JiraIssueLink notes a link between 2 Jira issues
type JiraIssueLink struct {
	// ID unique identifier of Jira Issue Link
	ID string

	// Type notes the class of issue link
	Type string

	// InID is the input Jira Issue ID
	InID string

	// OutID is the output Jira Issue ID
	OutID string

	// Comment on the jira issue link
	Comment *JiraIssueComment
}

// NewJiraIssueLink creates a new JiraIssueLink for a jira.IssueLink
func NewJiraIssueLink(from jira.IssueLink) JiraIssueLink {
	// Parse comment
	var com *JiraIssueComment = nil
	if from.Comment != nil {
		n := NewJiraIssueComment(*from.Comment)
		com = &n
	}

	return JiraIssueLink{
		ID:      from.ID,
		Type:    from.Type.Name,
		InID:    from.InwardIssue.ID,
		OutID:   from.OutwardIssue.ID,
		Comment: com,
	}
}

func (l JiraIssueLink) String() string {
	return fmt.Sprintf("ID: %s\n"+
		"Type: %s\n"+
		"InID: %s\n"+
		"OutID: %s\n"+
		"Comment: %s",
		l.ID, l.Type, l.InID, l.OutID, l.Comment)
}
