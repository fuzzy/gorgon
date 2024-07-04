package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/fuzzy/gorgon/cache"
	"github.com/fuzzy/gorgon/config"
	"github.com/fuzzy/gorgon/ghapi"
	"github.com/fuzzy/gorgon/utils"
)

func TestNewGorgonCache(t *testing.T) {
	dir := "test_cache_dir"
	defer os.RemoveAll(dir) // Clean up after test

	cache := cache.NewGorgonCache(dir)

	if cache.Dir != dir {
		t.Errorf("Expected Dir to be %s, got %s", dir, cache.Dir)
	}

	if cache.TempDir != "temp" {
		t.Errorf("Expected TempDir to be 'temp', got %s", cache.TempDir)
	}

	if !utils.Exists(dir) {
		t.Errorf("Expected directory %s to be created", dir)
	}

	if !utils.Exists("temp") {
		t.Errorf("Expected temp directory to be created")
	}
}

func TestInitCheck(t *testing.T) {
	dir := "test_cache_dir"
	defer os.RemoveAll(dir) // Clean up after test

	cache := &cache.GorgonCache{
		Dir:     dir,
		TempDir: "temp",
	}

	cache.InitCheck()

	if !utils.Exists(dir) {
		t.Errorf("Expected directory %s to be created", dir)
	}

	if !utils.Exists("temp") {
		t.Errorf("Expected temp directory to be created")
	}
}

func TestNewGorgonConfigDefault(t *testing.T) {
	config, err := config.NewGorgonConfig("nonexistent_file.json")
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

	config, err := config.NewGorgonConfig(filename)
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

	cfg := &config.GorgonConfig{
		Filename: filename,
		Username: "testuser",
		Projects: []config.GorgonProjectConfig{{ID: 1, Name: "testproject", Owner: "testowner"}},
		Repos:    []config.GorgonRepoConfig{{Name: "testrepo"}},
		Cache: config.GorgonCacheConfig{
			EnableJson: true,
			EnableYaml: false,
			Dir:        "/tmp/cache",
			TempDir:    "/tmp/cache/temp",
		},
		Agenda: config.CorgonAgendaConfig{
			OutputFile:  "/tmp/agenda.org",
			StatusMap:   map[string]string{"TODO": "To Do"},
			PriorityMap: map[string]string{"A": "High"},
		},
	}

	err := cfg.Save()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read saved config file: %v", err)
	}

	var loadedConfig config.GorgonConfig
	err = json.Unmarshal(data, &loadedConfig)
	if err != nil {
		t.Fatalf("Failed to unmarshal saved config: %v", err)
	}

	if loadedConfig.Username != "testuser" {
		t.Errorf("Expected Username to be testuser, got %s", loadedConfig.Username)
	}
}

// TestSyncDownStreamMock tests the SyncDownStreamMock function.
func TestSyncDownStreamMock(t *testing.T) {
	data := `{
		"items": [{
			"assignees": ["user1", "user2"],
			"content": {
				"body": "Sample body",
				"number": 1,
				"title": "Sample title",
				"url": "http://example.com",
				"type": "issue",
				"repository": "sample-repo"
			},
			"env": "production",
			"estimate - Hours": 5,
			"id": "item1",
			"labels": ["bug", "urgent"],
			"linked pull requests": ["pr1", "pr2"],
			"milestone": {
				"dueOn": "2022-01-01",
				"title": "Milestone 1",
				"description": "Sample milestone",
				"dueDate": "2022-01-01"
			},
			"priority": "high",
			"repository": "sample-repo",
			"size": 3,
			"sprint": {
				"duration": 14,
				"startDate": "2022-01-01",
				"endDate": "2022-01-14",
				"title": "Sprint 1"
			},
			"status": "open",
			"title": "Sample issue"
		}],
		"totalCount": 1
	}`
	filename := "test_items.json"
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to write test items file: %v", err)
	}
	defer os.Remove(filename)

	items, err := ghapi.SyncDownStreamMock(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if items.TotalCount != 1 {
		t.Errorf("Expected TotalCount to be 1, got %d", items.TotalCount)
	}

	if len(items.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items.Items))
	}

	item := items.Items[0]
	if item.Content.Title != "Sample title" {
		t.Errorf("Expected item title to be 'Sample title', got %s", item.Content.Title)
	}

	if item.Priority != "high" {
		t.Errorf("Expected item priority to be 'high', got %s", item.Priority)
	}
}

// TestSyncDownStreamMockFileNotFound tests SyncDownStreamMock when the file is not found.
func TestSyncDownStreamMockFileNotFound(t *testing.T) {
	_, err := ghapi.SyncDownStreamMock("nonexistent_file.json")
	if err == nil {
		t.Errorf("Expected an error when file does not exist, got nil")
	}
}

// TestSyncDownStreamMockInvalidJSON tests SyncDownStreamMock when the JSON is invalid.
func TestSyncDownStreamMockInvalidJSON(t *testing.T) {
	data := `invalid json`
	filename := "test_invalid_items.json"
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to write test invalid items file: %v", err)
	}
	defer os.Remove(filename)

	_, err = ghapi.SyncDownStreamMock(filename)
	if err == nil {
		t.Errorf("Expected an error when JSON is invalid, got nil")
	}
}
