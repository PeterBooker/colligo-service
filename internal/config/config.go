package config

import (
	"os"
)

type Service struct {
	Key    string
	Secret string
}

// Config contains global application information
type Config struct {
	Name        string
	Version     string
	Commit      string
	Date        string
	Environment string
	WD          string
	Twitter     Service
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
	tw := Service{
		Key:    os.Getenv("TWITTER_KEY"),
		Secret: os.Getenv("TWITTER_SECRET"),
	}
	c.Twitter = tw

	return c
}
