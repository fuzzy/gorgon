package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNewGorgonConfigDefault(t *testing.T) {
	config, err := NewGorgonConfig("nonexistent_file.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.Username != os.Getenv("USER") {
		t.Errorf("Expected Username to be %s, got %s", os.Getenv("USER"), config.Username)
	}

	home := os.Getenv("HOME")
	if config.Cache.Dir != filepath.Join(home, ".cache/gorgon") {
		t.Errorf("Expected Cache Dir to be %s, got %s", filepath.Join(home, ".cache/gorgon"), config.Cache.Dir)
	}

	if config.Agenda.OutputFile != filepath.Join(home, ".org-agenda/gorgon.org") {
		t.Errorf("Expected Agenda OutputFile to be %s, got %s", filepath.Join(home, ".org-agenda/gorgon.org"), config.Agenda.OutputFile)
	}
}

func TestGorgonConfigLoad(t *testing.T) {
	data := `{
		"username": "testuser",
		"projects": [{"id": 1, "name": "testproject", "owner": "testowner"}],
		"repos": [{"name": "testrepo"}],
		"cache": {
			"enableJson": true,
			"enableYaml": false,
			"dir": "/tmp/cache",
			"tempDir": "/tmp/cache/temp"
		},
		"agenda": {
			"outputFile": "/tmp/agenda.org",
			"statusMap": {"TODO": "To Do"},
			"priorityMap": {"A": "High"}
		}
	}`
	filename := "test_config.json"
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}
	defer os.Remove(filename)

	config, err := NewGorgonConfig(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.Username != "testuser" {
		t.Errorf("Expected Username to be testuser, got %s", config.Username)
	}

	if config.Projects[0].Name != "testproject" {
		t.Errorf("Expected Project Name to be testproject, got %s", config.Projects[0].Name)
	}
}

func TestGorgonConfigSave(t *testing.T) {
	filename := "test_save_config.json"
	defer os.Remove(filename)

	config := &GorgonConfig{
		filename: filename,
		Username: "testuser",
		Projects: []GorgonProjectConfig{{ID: 1, Name: "testproject", Owner: "testowner"}},
		Repos:    []GorgonRepoConfig{{Name: "testrepo"}},
		Cache: GorgonCacheConfig{
			EnableJson: true,
			EnableYaml: false,
			Dir:        "/tmp/cache",
			TempDir:    "/tmp/cache/temp",
		},
		Agenda: CorgonAgendaConfig{
			OutputFile:  "/tmp/agenda.org",
			StatusMap:   map[string]string{"TODO": "To Do"},
			PriorityMap: map[string]string{"A": "High"},
		},
	}

	err := config.Save()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read saved config file: %v", err)
	}

	var loadedConfig GorgonConfig
	err = json.Unmarshal(data, &loadedConfig)
	if err != nil {
		t.Fatalf("Failed to unmarshal saved config: %v", err)
	}

	if loadedConfig.Username != "testuser" {
		t.Errorf("Expected Username to be testuser, got %s", loadedConfig.Username)
	}
}
