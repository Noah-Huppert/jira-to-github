package models

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

// JiraIssueComment is a comment on a Jira Issue
type JiraIssueComment struct {
	// ID is the Jira Issue Comment unique ID
	ID string

	// AuthorKey is the Jira User's Key who wrote the comment
	AuthorKey string

	// Text is the Jira Issue Comment body
	Text string
}

// NewJiraIssueComment creates a new JiraIssueComment from a jira.Comment
func NewJiraIssueComment(from jira.Comment) JiraIssueComment {
	return JiraIssueComment{
		ID:        from.ID,
		AuthorKey: from.Author.Key,
		Text:      from.Body,
	}
}

func (c JiraIssueComment) String() string {
	return fmt.Sprintf("ID: %s\n"+
		"AuthorKey: %s\n"+
		"Text: %s",
		c.ID, c.AuthorKey, c.Text)
}
