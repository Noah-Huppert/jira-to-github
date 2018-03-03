package aggr

import (
	"fmt"
	"strconv"

	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/Noah-Huppert/jira-to-github/store"
)

// JiraAggregateStore indicates the store name to save JiraAggregates into
const JiraAggregateStore string = "jira_aggregates"

// JiraAggregateStoreKey is the key Jira Aggregates will always be stored under
const JiraAggregateStoreKey string = "main"

// JiraAggregate provides an overview of all the retrieved Jira issues
type JiraAggregate struct {
	// IDs is a list of all the Jira Issue IDs
	IDs []int

	// MaxID indicates the highest Jira Issue ID we have
	MaxID int

	// UserKeys holds all the Jira User keys present in the retrieved issues

	// aggregated indicates if the data has been populated
	aggregated bool
}

// NewJiraAggregate creates a new JiraAggregate
func NewJiraAggregate() *JiraAggregate {
	return &JiraAggregate{
		aggregated: false,
	}
}

func (a JiraAggregate) String() string {
	return fmt.Sprintf("IDs: %d\n"+
		"MaxID: %d\n"+
		"aggregated: %t",
		a.IDs, a.MaxID, a.aggregated)
}

// Aggregate collects the information. Pass the Jira Issue Store.
func (a *JiraAggregate) Aggregate(stores *store.Stores) error {
	// Get all Jira issue keys
	keys, err := stores.Jira.Issues.Keys()
	if err != nil {
		return fmt.Errorf("error retrieving all Jira Issue keys: %s",
			err.Error())
	}

	// Get each issue
	for _, key := range keys {
		var issue models.JiraIssue = models.JiraIssue{}

		err := stores.Jira.Issues.Get(key, &issue)
		if err != nil {
			return fmt.Errorf("error retrieving key: %s, err: %s",
				key, err.Error())
		}

		// Add to IDs
		id, err := strconv.Atoi(issue.ID)
		if err != nil {
			return fmt.Errorf("error converting issue id to int "+
				"issue: %s, err: %s", issue, err.Error())
		}
		a.IDs = append(a.IDs, id)

		// Check if greater than max
		if id > a.MaxID {
			a.MaxID = id
		}
	}

	a.aggregated = true

	return nil
}
