package config

import (
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
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

	selfContID := getSelfContainerID()

	if selfContID != "" {
		log.Debugf("Self config id: %s", selfContID)

		env.CurrentContainerID = selfContID
	}

	log.Debugf("Environment loaded: %v", env)

	return &env
}

func getSelfContainerID() string {
	hostNameCmd := exec.Command("cat", "/etc/hostname")

	selfContID, err := hostNameCmd.Output()

	if err != nil {
		log.Warnf("Failed to cat /etc/hostname: %s", err)
	}

	return strings.TrimSpace(string(selfContID))
}
