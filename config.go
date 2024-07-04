package main

import (
	"encoding/json"
	"os"
)

type GorgonConfig struct {
	ProjectName string `json:"projectName"`
	ProjectID   int    `json:"projectID"`
	Owner       string `json:"owner"`
	Username    string `json:"username"`

	CacheDir  string `json:"cacheDir"`
	AgendaDir string `json:"agendaDir"`
}

func NewGorgonConfig(fn string) (*GorgonConfig, error) {
	retv := &GorgonConfig{}

	_, err := os.Stat(fn)
	if err != nil {
		return nil, err
	}

	retv.load(fn)

	return retv, nil
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
