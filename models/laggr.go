package models

import (
	"encoding/json"
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/hash"
	"github.com/fatih/structs"
)

// LinkAggregateStore is the store used to save link aggregates
const LinkAggregateStore string = "link_aggregates"

// LinkAggregate collects all links for a specific model
type LinkAggregate struct {
	// Model is the model the aggregate is generated for
	Model string

	// JiraIndex maps Jira IDs to GH IDs
	JiraIndex map[string]string
}

// NewLinkAggregate creates a new LinkAggregate instance
func NewLinkAggregate() *LinkAggregate {
	return &LinkAggregate{
		JiraIndex: map[string]string{},
	}
}

func (a LinkAggregate) String() string {
	return fmt.Sprintf("Model: %s\n"+
		"JiraIndex: %s",
		a.Model, a.JiraIndex)
}

// Hash implements hash.Hashable.Hash
func (a LinkAggregate) Hash() string {
	return hash.HashStr(a.String())
}

// MarshalJSON implements json.Marshaler.MarshalJSON
func (a LinkAggregate) MarshalJSON() ([]byte, error) {
	m := structs.Map(a)

	m["hash"] = a.Hash()

	return json.Marshal(m)
}
