package main

import (
	"fmt"
	"io"
	"os"

	"github.com/niklasfasching/go-org/org"
)

func parseOrgFile(fname string) error {
	r, p := io.Reader(nil), fname
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	r = f

	// Parse the org file
	d := org.New().Parse(r, p)
	if d == nil {
		return fmt.Errorf("Failed to parse org file %s", p)
	}
	fmt.Printf("Parsed %d nodes\n", len(d.Nodes))

	return nil
}
