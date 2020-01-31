package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " ")) // strings.join , concatenates all stuff inside terms string slice, QueryEscape escapes ? and &
	resp, err := http.Get(IssuesURL + "?q=" + q)   // Executing GET on IssuesURL with the supplied terms
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Failed for some reason %s\n", resp.StatusCode)
	}

	var result IssuesSearchResult

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
