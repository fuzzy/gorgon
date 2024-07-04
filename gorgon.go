package main

import (
	"fmt"
	"os"

	"github.com/fuzzy/gorgon/config"
	"github.com/fuzzy/gorgon/ghapi"
)

func main() {
	opts := parseArgs()
	cnfg, err := config.NewGorgonConfig(opts.Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Config: %s\n", opts.Config)
	fmt.Printf("Username: %s\n", cnfg.Username)

	cmd := fmt.Sprintf("gh project item-list %d --owner %s -L %d --format json", opts.ID, opts.Owner, opts.Results)
	fmt.Println(cmd)

	items, err := ghapi.SyncDownStreamMock("output.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("JSON Items: %d\n", items.TotalCount)
	fmt.Println("JSON Parsing Worked!")

	parseOrgFile("github.org")
	fmt.Println("Org Parsing Worked!")
}
