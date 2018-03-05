# Jira To GitHub
Migrates Jira issues to Github
 
# Table Of Contents
- [Overview](#overview)
- [Install](#install)
- [Setup](#setup)
- [Run](#run)

# Overview
Jira to GitHub is a command line tool for migrating Jira issues to GitHub.  

It is built from the ground up to ensure a smooth migration from Jira to 
GitHub. Jira to GitHub will never modify existing GitHub issues. Nor will it 
ever create duplicates.

See the [Setup](#setup) and [Running](#running) sections for more information.

# Install
To install Jira To GitHub you must build from source. In the future binary 
distributions may be provided.  

First clone down this repository with `go get github.com/Noah-Huppert/jira-to-github`.  
Then run `make install`.  

The Jira to GitHub migration tool will then be installed in `$GOPATH/bin` 
with the name `j2gh`.

# Setup
To run Jira To GitHub you must edit the configuration file.  

The configuration file holds the following values:

- Jira: Jira specific configuration
	- URL: Location of Jira instance (include scheme in URL)
	- Project: Jira project to extract issues from
	- Username: Jira account username to authenticate with
	- Password: Jira account password
- GitHub : GitHub related configuration
	- AccessToken: GitHub API access token used to authenticate with the 
		       GitHub API. 
		- Must have the following permissions `repo:public_repo`, 
		  `admin:org:read:org`, `user:read:user`, `user:user:email`, 
		  `write:discussion:read:discussion`.
	- RepoOwner: GitHub user / organization who owns repository to migrate 
		     Jira issues to. 
	- RepoName: Name of GitHub repository to migrate Jira issues to.

First copy `config.example.toml` to `config.toml`.  
Then modify it with your own values.  

# Run
To run Jira to GitHub simply run the executable which is named `j2gh`.  

Usage: `j2gh <Command>`.  

There are 3 commands:

- `fetch`: Will retrieve state from the Jira and GitHub APIs
	- Takes no arguments.
- `link <Model> <Jira ID> <GitHub ID>`: Will create an association between a 
	Jira and GitHub API entity
	- `Model`: The type of model to create a link for. Valid values are 
		`user`, `label`, `issue`.
	- `Jira ID`: ID of Jira model
	- `GitHub ID`: ID of GitHub model
	- If run with no arguments the `link` command lists all the models 
	  that can be linked, and their link status.
- `create`: Will create the GitHub issues for the retrieved Jira issues
	- Takes no arguments.

Jira to GitHub is safe to run multiple times. As it is aware of the issues 
which it has already transfered over.
