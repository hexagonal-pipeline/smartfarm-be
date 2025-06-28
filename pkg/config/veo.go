package config

import (
	"os"

	"github.com/samber/do/v2"
)

type VeoConfig struct {
	ProjectID string
	Location  string
}

func NewVeoConfig(_ do.Injector) (*VeoConfig, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = "smartfarm-project" // fallback
	}

	location := os.Getenv("GOOGLE_CLOUD_LOCATION")
	if location == "" {
		location = "us-central1" // default location
	}

	return &VeoConfig{
		ProjectID: projectID,
		Location:  location,
	}, nil
}
