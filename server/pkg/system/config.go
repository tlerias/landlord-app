package system

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	// ConfigSpecification holds configurations
	ConfigSpecification struct {
		// Default application values
		AppName     string        `envconfig:"APP_NAME" default:"Placeholder"`
		Debug       bool          `envconfig:"DEBUG" default:"true"`
		LogLevel    string        `envconfig:"LOG_LEVEL" default:"info"`
		CacheTTL    time.Duration `envconfig:"CACHE_TTL" default:"10m"`
		Environment string        `envconfig:"ENVIRONMENT" default:"local"`
	}
)

// LoadEnvironmentConfig loads environment configuration variables
func LoadEnvironmentConfig() (*ConfigSpecification, error) {
	var config ConfigSpecification

	err := envconfig.Process("", &config)
	if err != nil {
		err = fmt.Errorf("LoadEnv: %s", err)
	}
	return &config, err
}
