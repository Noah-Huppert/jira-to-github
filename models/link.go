package models

import (
	"encoding/json"
	"fmt"

	"github.com/Noah-Huppert/jira-to-github/hash"
	"github.com/fatih/structs"
)

// Link indicates an association between a Jira and GitHub entity. The type of
// entity will be determined by the name of the store the Link is saved in.
type Link struct {
	// GitHubID is the ID of the GitHub entity which is linked
	GitHubID string

	// GitHubHash is the hash of the GitHub entity at the time of linking.
	GitHubHash string

	// JiraID is the ID of the Jira entity which is linked
	JiraID string

	// JiraHash is the hash of the Jira entity at the time of linking.
	JiraHash string
}

func (l Link) String() string {
	return fmt.Sprintf("GitHubID: %d\n"+
		"GitHubHash: %s\n"+
		"JiraID: %s\n"+
		"JiraHash: %s",
		l.GitHubID, l.GitHubHash, l.JiraID, l.JiraHash)
}

// Hash implements hash.Hashable.Hash
func (l Link) Hash() string {
	return hash.HashStr(l.String())
}

// MarshalJSON implements json.Marshalable.Marshal
func (l Link) MarshalJSON() ([]byte, error) {
	m := structs.Map(l)

	m["hash"] = l.Hash()

	return json.Marshal(m)
}
