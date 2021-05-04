package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var AppConfig Config

func InitialiseConfig() error {
	err := envconfig.Process("server", &AppConfig)
	if err != nil {
		return fmt.Errorf("InitialiseConfig: %v", err)
	}
	return nil
}
