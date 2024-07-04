package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// GorgonCacheConfig is a struct that holds the configuration for the cache
// directory and the cache format options.
type GorgonCacheConfig struct {
	// Enable JSON caching, defaults to true
	EnableJson bool `json:"enableJson"`
	// Enable YAML caching, defaults to false
	EnableYaml bool `json:"enableYaml"`
	// Cache directory
	Dir string `json:"dir"`
	// Last sync directory
	TempDir string `json:"tempDir"`
}

// CorgonAgendaConfig is a struct that holds the configuration for the agenda
// output file, status map, and priority map.
type CorgonAgendaConfig struct {
	// Output file for the agenda
	OutputFile string `json:"outputFile" `
	// Status map for the agenda
	StatusMap map[string]string `json:"statusMap"`
	// Priority map for the agenda
	PriorityMap map[string]string `json:"priorityMap"`
}

// GorgonProjectConfig is a struct that holds the configuration for a GitHub
// project to sync issues with.
type GorgonProjectConfig struct {
	// Project ID
	ID int `json:"id"`
	// Project Name
	Name string `json:"name"`
	// Project Owner
	Owner string `json:"owner"`
}

// GorgonRepoConfig is a struct that holds the configuration for a GitHub
// repository to sync issues with.
type GorgonRepoConfig struct {
	// Repo Name
	Name string `json:"name"`
}

type GorgonConfig struct {
	filename string
	// GitHub User (this should be the gorgon user's GitHub username)
	// We do no authentication setup in this project, that is handeled
	// by the user's GitHub CLI setup, so we just need the username for
	// gh cli arguments, and processing.
	Username string `json:"username"`
	// Github projects to sync
	Projects []GorgonProjectConfig `json:"projects"`
	// GitHub Repos to sync
	Repos []GorgonRepoConfig `json:"repos"`
	// Cache configuration
	Cache GorgonCacheConfig `json:"cache"`
	// Agenda configuration (tailor to your org configuration)
	Agenda CorgonAgendaConfig `json:"agenda"`
}

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
