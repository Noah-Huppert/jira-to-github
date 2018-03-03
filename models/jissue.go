package models

import (
	"encoding/json"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"strings"
	//	"time"

	"github.com/Noah-Huppert/jira-to-github/hash"
	"github.com/Noah-Huppert/jira-to-github/str"
	"github.com/fatih/structs"
)

// JiraDateFormat parses a Jira date in the format
//                            "2018-03-01T16:45:05.866-0500"
//const JiraDateFormat string = "2006-01-02T15:04:05.999999999Z07-0700"

// JiraIssueStore is the name of the store to keep Jira Issues in
const JiraIssueStore string = "jira_issues"

// JiraIssue holds the information we can tranfer to GitHub from a Jira issue.
type JiraIssue struct {
	// ID holds the Jira issue ID.
	ID string

	// Type the style of Jira issue.
	Type string

	// ProjectID holds the ID of the Jira project which the issue belongs to.
	ProjectID string

	// AssigneeKey holds the Key of the Jira user who is assigned to the
	// issue
	AssigneeKey string

	// Assignee holds the JiraUser model for the corresponding
	// AsigneeKey. Only populated when the JiraIssue is created from a
	// Jira API query response.
	Assignee *JiraUser `json:"-"`

	// Resolution holds the final status of the Jira issue
	Resolution string

	// Priority indicate how urgent an issue is
	Priority string

	// Status indicates the current work state of the issue
	Status string

	// Title holds a short summary of the Jira issue
	Title string

	// Description holds a longer write up about the Jira issue
	Description string

	// Progress indicates how complete an issue is. In the range from:
	// [0, 1]
	Progress float32

	// Links holds the Jira Issue Links associated with the issue
	Links []JiraIssueLink

	// Comments holds the Jira Issue Comments associated with the issue
	Comments []JiraIssueComment

	// Labels holds the Jira labels put in the issue
	Labels []string

	// Updated holds the last time the Jira issue was updated
	//Updated *time.Time
}

// NewJiraIssue creates a new JiraIssue from a jira.Issue. An error is returned
// if one occurs.
func NewJiraIssue(from jira.Issue) (JiraIssue, error) {
	// Parse assignee
	assnKey := ""
	var assn *JiraUser = nil

	if from.Fields.Assignee != nil {
		assnKey = from.Fields.Assignee.Key

		user := NewJiraUser(*from.Fields.Assignee)
		assn = &user
	}

	// Parse resolution
	res := ""
	if from.Fields.Resolution != nil {
		res = from.Fields.Resolution.Name
	}

	// Parse progress
	var prog float32

	if from.Fields.Progress.Total != 0 {
		prog = float32(from.Fields.Progress.Progress) / float32(from.Fields.Progress.Total)
	}

	// Parse links
	links := []JiraIssueLink{}
	for _, jLink := range from.Fields.IssueLinks {
		links = append(links, NewJiraIssueLink(*jLink))
	}

	// Parse comments
	comments := []JiraIssueComment{}

	if from.Fields.Comments != nil {
		for _, jCom := range from.Fields.Comments.Comments {
			comments = append(comments, NewJiraIssueComment(*jCom))
		}
	}

	// Parse updated
	/*var updated *time.Time = nil

	t, err := time.Parse(JiraDateFormat, from.Fields.Updated)
	if err != nil {
		return JiraIssue{}, fmt.Errorf("error parsing issue updated "+
			"field into date: %s", err.Error())
	}
	updated = &t*/

	return JiraIssue{
		ID:          from.ID,
		Type:        from.Fields.Type.Name,
		AssigneeKey: assnKey,
		Assignee:    assn,
		ProjectID:   from.Fields.Project.ID,
		Resolution:  res,
		Priority:    from.Fields.Priority.Name,
		Status:      from.Fields.Status.Name,
		Title:       from.Fields.Summary,
		Description: from.Fields.Description,
		Progress:    prog,
		Links:       links,
		Comments:    comments,
		Labels:      from.Fields.Labels,
		//Updated:     updated,
	}, nil
}

// String implements fmt.Stringer.String
func (i JiraIssue) String() string {
	// TODO Figure out why slices not casting to slice of Stringers
	return fmt.Sprintf("ID: %s\n"+
		"Type: %s\n"+
		"ProjectID: %s\n"+
		"AssigneeKey: %s\n"+
		"Resolution: %s\n"+
		"Priority: %s\n"+
		"Status: %s\n"+
		"Title: %s\n"+
		"Description: %s\n"+
		"Progress: %g\n"+
		"Links: %s\n"+
		"Comments: %s\n"+
		"Labels: [%s]",
		//"Updated: %s",
		i.ID, i.Type, i.ProjectID, i.AssigneeKey, i.Resolution,
		i.Priority, i.Status, i.Title, i.Description, i.Progress,
		str.JoinStringers(StringerFromJILinks(i.Links)),
		str.JoinStringers(StringerFromJIComments(i.Comments)),
		strings.Join(i.Labels, ", "))
	//,i.Updated)
}

// Hash implements Hashable.Hash
func (i JiraIssue) Hash() string {
	return hash.HashStr(i.String())
}

// MarshalJSON implements json.Marshaler.MarshalJSON
func (i JiraIssue) MarshalJSON() ([]byte, error) {
	m := structs.Map(i)

	m["hash"] = i.Hash()

	return json.Marshal(m)
}
