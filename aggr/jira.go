package aggr

import (
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/store"
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

// Aggregate collects the information. An error will be returned if on occurs
func (a *JiraAggregate) Aggregate(stores *store.Stores) error {
	// Issues
	if err := a.Issues.Aggregate(stores); err != nil {
		return fmt.Errorf("error aggregating issues: %s", err.Error())
	}

	a.aggregated = true

	return nil
}
