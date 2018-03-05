package config

// GitHub holds GitHub configuration
type GitHub struct {
	// AccessToken is the personal access token used to authenticate with
	// GitHub.
	AccessToken string

	// RepoOwner is the owner of the GitHub repository to migrate issues to.
	RepoOwner string

	// RepoName is the name of the GitHub repository to migrate issues to.
	RepoName string
}
