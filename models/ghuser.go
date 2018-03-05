package models

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/google/go-github/github"

	"github.com/Noah-Huppert/jira-to-github/hash"
)

// GitHubUserStore holds the name of the store used to save GitHub user
// models.
const GitHubUserStore string = "github_users"

// GitHubUser holds information about a GitHub user.
type GitHubUser struct {
	// ID is the User's unique identifier.
	ID int64

	// Login is the User's username
	Login string

	// Email is the User's email address
	Email string

	// Name is the User's name
	Name string
}

// NewGitHubUser creates a new GitHubUser from a github.User
func NewGitHubUser(from github.User) GitHubUser {
	return GitHubUser{
		ID:    *from.ID,
		Login: *from.Login,
		Email: *from.Email,
		Name:  *from.Name,
	}
}

func (u GitHubUser) String() string {
	return fmt.Sprintf("ID: %s\n"+
		"Login: %s\n"+
		"Email: %s\n"+
		"Name: %s",
		u.ID, u.Login, u.Email, u.Name)
}

// MarshalJSON implements json.Marshaler.MarshalJSON
func (u GitHubUser) MarshalJSON() ([]byte, error) {
	m := structs.Map(u)

	m["hash"] = u.Hash()

	return json.Marshal(u)
}

// Hash implements hash.Hashable.Hash
func (u GitHubUser) Hash() string {
	return hash.HashStr(u.String())
}
