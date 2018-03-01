package config

// Jira holds Jira specific configuration
type Jira struct {
	// URL is the URL of the Jira instance to retrieve information from.
	// Must be a full URL, with a scheme and host.
	URL string

	// Project is the Jira project to retrieve issues for
	Project string

	// Username is the name of the user account to authenticate with
	Username string

	// Password is the user password value to authenticate with
	Password string
}
