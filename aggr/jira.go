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

	// Labels holds all the labels present in the retrieved issues
	Labels []string

	// aggregated indicates if the data has been populated
	aggregated bool
}

// NewJiraAggregate creates a new JiraAggregate
func NewJiraAggregate() *JiraAggregate {
	return &JiraAggregate{
		Labels:     []string{},
		aggregated: false,
	}
}

func (a JiraAggregate) String() string {
	return fmt.Sprintf("IDs: %d\n"+
		"MaxID: %d\n"+
		"Labels: %s\n"+
		"aggregated: %t",
		a.IDs, a.MaxID, a.Labels, a.aggregated)
}

// Aggregate collects the information. Pass the Jira Issue Store.
func (a *JiraAggregate) Aggregate(stores *store.Stores) error {
	// Get all Jira issue keys
	keys, err := stores.Jira.Issues.Keys()
	if err != nil {
		return fmt.Errorf("error retrieving all Jira Issue keys: %s",
			err.Error())
	}

	// tempLabels is a map which temporarily records which labels the
	// issues have
	tempLabels := map[string]bool{}

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

		// Record labels
		for _, label := range issue.Labels {
			tempLabels[label] = true
		}
	}

	// Save labels
	for label := range tempLabels {
		a.Labels = append(a.Labels, label)
	}

	a.aggregated = true

	return nil
}
