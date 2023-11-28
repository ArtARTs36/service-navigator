package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Environment struct {
	User               string `default:"developer"`
	CurrentContainerID string
}

func InitEnvironment() *Environment {
	var env Environment

	log.Debug("Loading environment")

	err := envconfig.Process("", &env)

	if err != nil {
		panic(fmt.Sprintf("Failed to load environment: %s", err))
	}

	selfContID, err := os.Hostname()

	if selfContID != "" {
		log.Debugf("Self config id: %s", selfContID)

		env.CurrentContainerID = selfContID
	} else if err != nil {
		log.Warnf("failed to fetch hostname: %s", err)
	}

	log.Debugf("Environment loaded: %v", env)

	return &env
}
