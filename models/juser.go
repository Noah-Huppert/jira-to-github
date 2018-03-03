package models

import (
	"encoding/json"
	"fmt"
	"github.com/Noah-Huppert/jira-to-github/hash"
	"github.com/andygrunwald/go-jira"
	"github.com/fatih/structs"
)

// JiraUserStore is the name of the store used to saved JiraUser models
const JiraUserStore string = "jira_users"

// JiraUser holds information about a Jira user
type JiraUser struct {
	// Key is a unique user handle
	Key string

	// Name is the user's name
	Name string

	// Email is the user's email address
	Email string

	// DisplayName is the name the user chooses to display
	DisplayName string
}

// NewJiraUser creates a new JiraUser from a jira.User
func NewJiraUser(from jira.User) JiraUser {
	return JiraUser{
		Key:         from.Key,
		Name:        from.Name,
		Email:       from.EmailAddress,
		DisplayName: from.DisplayName,
	}
}

func (u JiraUser) String() string {
	return fmt.Sprintf("Key: %s\n"+
		"Name: %s\n"+
		"Email: %s\n"+
		"DisplayName: %s",
		u.Key, u.Name, u.Email, u.DisplayName)
}

// MarshalJSON implements json.Marshaler.MarshalJSON
func (u JiraUser) MarshalJSON() ([]byte, error) {
	m := structs.Map(u)

	m["hash"] = u.Hash()

	return json.Marshal(m)
}

// Hash implements hash.Hashable.Hash
func (u JiraUser) Hash() string {
	return hash.HashStr(u.String())
}
