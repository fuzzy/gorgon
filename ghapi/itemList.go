package ghapi

import (
	"encoding/json"
	"fmt"
	"os"
)

// ItemContent is a struct that holds the body, number, title, URL, type, and repository.
// NOTE: Add more items as found from the JSON output of the gh project item-list command.
type ItemContent struct {
	Body       string `json:"body"`
	Number     int    `json:"number"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	Repository string `json:"repository"`
}

// ItemSprint is a struct that holds the duration, start and end dates, and title of a sprint.
// NOTE: Add more items as found from the JSON output of the gh project item-list command.
type ItemSprint struct {
	Duration  int    `json:"duration"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Title     string `json:"title"`
}

// ItemMilestone is a struct that holds the due date, title, due date, and description of a milestone.
// NOTE: Add more items as found from the JSON output of the gh project item-list command.
type ItemMilestone struct {
	DueOn       string `json:"dueOn"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
}

// Item is a struct that holds the assignees, content, environment, estimate, ID, labels,
// linked pull requests, milestone, priority, repository, size, sprint, status, and title.
// NOTE: Add more items as found from the JSON output of the gh project item-list command.
type Item struct {
	Assignees          []string      `json:"assignees"`
	Content            ItemContent   `json:"content"`
	Env                string        `json:"env"`
	Estimate           float64       `json:"estimate - Hours"`
	Id                 string        `json:"id"`
	Labels             []string      `json:"labels"`
	LinkedPullRequests []string      `json:"linked pull requests"`
	Milestone          ItemMilestone `json:"milestone"`
	Priority           string        `json:"priority"`
	Repository         string        `json:"repository"`
	Size               int           `json:"size"`
	Sprint             ItemSprint    `json:"sprint"`
	Status             string        `json:"status"`
	Title              string        `json:"title"`
}

// ItemList is a struct that holds the items and the total count. This is the top-level struct
// as returned by the gh project item-list command.
type ItemList struct {
	Items      []Item `json:"items"`
	TotalCount int    `json:"totalCount"`
}

// SyncDownStream is a function that takes a command string and returns an ItemList and an error.
func SyncDownStream(cmd string) (*ItemList, error) {
	fmt.Println(cmd)

	return nil, nil
}

// This function is a placeholder for the actual implementation of the GitHub API call. It is used
// to mock the API call for testing purposes. Please do not use this function for anything.
func SyncDownStreamMock(fname string) (*ItemList, error) {
	retv := ItemList{}
	buf, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, &retv); err != nil {
		return nil, err
	}

	return &retv, nil
}
