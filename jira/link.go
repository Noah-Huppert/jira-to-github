package jira

import (
	"fmt"
	"os"

	"github.com/Noah-Huppert/jira-to-github/models"
	"github.com/Noah-Huppert/jira-to-github/store"
)

// UpdateLinkAggregates updates the link aggregates for the user, label, and
// issue models. An error will be returned if one occurs.
func UpdateLinkAggregates(stores *store.Stores) error {
	// models are the names of the models which can have links
	models := []string{models.UserModelName, models.IssueModelName,
		models.LabelModelName}
	jStores := []store.Store{stores.Jira.Users, stores.Jira.Issues}

	for _, model := range models {

	}

	return nil
}

// UpdateSingleLinksAggregate updates and stores a LinkAggregate.
// Args:
//	- model: Name of model we are aggregating links for
//		- Should be models.UserModelName, models.IssueModelName,
//		  models.LabelModelName
//	- jStore: Jira model store to retrieve models to find links for
//	- lStore: Links store used to save links for model
//
// An error will be returned if one occurs.
func UpdateSingleLinksAggregate(model string, jStore store.Store,
	lStore store.LinkAggregateStore, stores *store.Stores) error {
	a := models.NewLinkAggregate()

	// Find link for each jira model
	keys, err := jStore.Keys()
	if err != nil {
		return fmt.Errorf("error retrieving Jira model keys"+
			", model: %s, err: %s",
			model, err.Error())
	}

	for _, key := range keys {
		// Find link for model
		var aggr models.LinkAggregate = models.LinkAggregate{}
		err = lStore.Get(key, &link)

		// If no link for model
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return fmt.Errorf("error retrieving link for model"+
				", model: %s, key: %s, err: %s",
				model, key, err.Error())
		}

		// TODO: Fix commented out syntax error below
		//a.JiraIndex[key] = .GitHubID
	}

	// Save
	// Under id: plural form of model arg
	id := fmt.Sprintf("%ss", model)
	if err = stores.Aggregates.Links.Set(id, a); err != nil {
		return fmt.Errorf("error saving link aggregate, model: %s"+
			", err: %s", model, err.Error())
	}

	return nil
}
