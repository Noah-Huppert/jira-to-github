package models

import (
	"encoding/json"
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/hash"
	"github.com/fatih/structs"
)

// JiraAggregateStore indicates the store name to save JiraAggregates into
const JiraAggregateStore string = "jira_aggregates"

// JiraAggregateStoreKey is the key Jira Aggregates will always be stored under
const JiraAggregateStoreKey string = "main"

// JiraAggregate provides an overview of all the retrieved Jira models
type JiraAggregate struct {
	// Issues holds the JiraIssuesAggregate
	Issues *JiraIssuesAggregate

	// aggregated indicates if the data has been populated
	aggregated bool
}

// NewJiraAggregate creates a new JiraAggregate
func NewJiraAggregate() *JiraAggregate {
	// Issues
	issues := NewJiraIssuesAggregate()

	return &JiraAggregate{
		Issues:     issues,
		aggregated: false,
	}
}

func (a JiraAggregate) String() string {
	return fmt.Sprintf("Issues: {%s}\n"+
		"aggregated: %t",
		a.Issues, a.aggregated)
}

// Hash implements hash.Hashable.Hash
func (a JiraAggregate) Hash() string {
	return hash.HashStr(a.String())
}

// MarshalJSON implements json.Marshaler.MarshalJSON
func (a JiraAggregate) MarshalJSON() ([]byte, error) {
	m := structs.Map(a)

	m["hash"] = a.Hash()

	return json.Marshal(m)
}
