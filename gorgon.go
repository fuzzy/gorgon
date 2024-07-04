package main

import (
	"fmt"
	"os"
)

func main() {
	opts := parseArgs()

	cmd := fmt.Sprintf("gh project item-list %d --owner %s -L %d --format json", opts.ID, opts.Owner, opts.Results)
	fmt.Println(cmd)

	items, err := syncDownStreamMock("output.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("JSON Items: %d\n", items.TotalCount)
	fmt.Println("JSON Parsing Worked!")

	parseOrgFile("github.org")
	fmt.Println("Org Parsing Worked!")
}
