package gthub

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	"time"
	"net/url"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
        TotalCount int `json:"total_count"`
        Items []*Issue
}

type Issue struct {
        Number int
        HTMLURL string `json:"html_url"`
        Title string
        State string
        User *User
        Created time.Time `json:"created_at"`
        Body string
}

type User struct {
        Login string
        HTMLURL string `json:"html_url"`
}


func SearchIssues(uls []string)(*IssueSearchResult, error) {
        q := url.QueryEscape(strings.Join(uls, " "))
        resp, err := http.Get(IssuesURL + "?q=" + q)
        if err != nil {
                return nil, err
        }

        if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("query failed: %s\n", resp.Status)
        }
        var result IssueSearchResult
        if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
                resp.Body.Close()
                return nil, err
        }
        resp.Body.Close()
        return &result, nil
}
