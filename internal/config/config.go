package config

import (
	"os"
)

// Config contains global application information
type Config struct {
	Name          string
	Version       string
	Commit        string
	Date          string
	Environment   string
	WD            string
	UpdateWorkers int
	SearchWorkers int
	Host          string
	Domains       string
	Standalone    bool
	DevMode       bool
}

// New creates a new Config from flags and environment variables
func New(version, commit, date, env string) *Config {
	c := &Config{
		Name:        "Colligo",
		Version:     version,
		Commit:      commit,
		Date:        date,
		Environment: env,
	}

	// Get Working Dir
	wd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	c.WD = wd

	// Get Env Vars
	// os.Getenv("FOO")

	return c
}
