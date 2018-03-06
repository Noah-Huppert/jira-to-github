package models

import (
	"fmt"
)

// JiraIssuesAggregate collects all the information related to Jira Issue models
type JiraIssuesAggregate struct {
	// IDs is a list of all the Jira Issue IDs
	IDs []int

	// MaxID indicates the highest Jira Issue ID we have
	MaxID int

	// Labels holds all the labels present in the retrieved issues
	Labels []string

	// aggregated indicates if the data has been populated
	aggregated bool
}

// NewJiraIssuesAggregate creates a new JiraIssuesAggregate
func NewJiraIssuesAggregate() *JiraIssuesAggregate {
	return &JiraIssuesAggregate{
		Labels:     []string{},
		aggregated: false,
	}
}

func (a JiraIssuesAggregate) String() string {
	return fmt.Sprintf("IDs: %d\n"+
		"MaxID: %d\n"+
		"Labels: %s\n"+
		"aggregated: %t",
		a.IDs, a.MaxID, a.Labels, a.aggregated)
}
