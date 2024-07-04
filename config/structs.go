package config

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
