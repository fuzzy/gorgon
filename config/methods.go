package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewGorgonConfig creates a new GorgonConfig object from the specified file.
// If the file does not exist, it will return an GorgonConfig object with
// default values set.
// If the file does exist, it will load the configuration from the file.
// If there is an error loading the file, it will return an error.
func NewGorgonConfig(fn string) (*GorgonConfig, error) {
	// Create a new GorgonConfig object
	retv := &GorgonConfig{
		filename: fn,
		Username: os.Getenv("USER"),
		Projects: []GorgonProjectConfig{},
		Repos:    []GorgonRepoConfig{},
		Cache: GorgonCacheConfig{
			EnableJson: true,
			EnableYaml: false,
			Dir:        fmt.Sprintf("%s/.cache/gorgon", os.Getenv("HOME")),
			TempDir:    fmt.Sprintf("%s/.cache/gorgon/temp", os.Getenv("HOME")),
		},
		Agenda: CorgonAgendaConfig{
			OutputFile:  fmt.Sprintf("%s/.org-agenda/gorgon.org", os.Getenv("HOME")),
			StatusMap:   map[string]string{},
			PriorityMap: map[string]string{},
		},
	}

	if _, err := os.Stat(fn); err != nil {
		return retv, nil
	}

	if err := retv.load(fn); err != nil {
		return nil, err
	}

	return retv, nil
}

// Save writes the GorgonConfig object to the file specified in the object.
// If there is an error writing the file, it will return an error.
func (gc *GorgonConfig) Save() error {
	data, err := json.MarshalIndent(gc, "", "  ")
	if err != nil {
		return err
	}

	_, err = os.Stat(fmt.Sprintf("%s/.config/gorgon", os.Getenv("HOME")))
	if err != nil {
		os.MkdirAll(fmt.Sprintf("%s/.config/gorgon", os.Getenv("HOME")), 0755)
	}

	ofp, err := os.Create(gc.filename)
	if err != nil {
		return err
	}
	defer ofp.Close()

	_, err = ofp.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (gc *GorgonConfig) load(fn string) error {
	// Load the config file
	data, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	// Unmarshal the data
	err = json.Unmarshal(data, gc)
	if err != nil {
		return err
	}

	return nil
}
