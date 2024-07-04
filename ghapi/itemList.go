package ghapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type ItemContent struct {
	Body       string `json:"body"`
	Number     int    `json:"number"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	Repository string `json:"repository"`
}

type ItemSprint struct {
	Duration  int    `json:"duration"`
	StartDate string `json:"startDate"`
	Title     string `json:"title"`
}

type ItemMilestone struct {
	DueOn       string `json:"dueOn"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

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

type ItemList struct {
	Items      []Item `json:"items"`
	TotalCount int    `json:"totalCount"`
}

func SyncDownStream(cmd string) (*ItemList, error) {
	fmt.Println(cmd)

	return nil, nil
}

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
