// Provides API for GitHub issues

package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Id		  int
	CreatedAt time.Time `json:"created_at`
	Body      string    // According to the book, this should contain the markdown body of the issue
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
