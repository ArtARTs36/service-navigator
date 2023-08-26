package config

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	User               string `default:"developer"`
	CurrentContainerID string
}

func InitEnvironment() *Environment {
	var env Environment

	log.Print("Loading environment")

	err := envconfig.Process("", &env)

	if err != nil {
		panic(fmt.Sprintf("Failed to load environment: %s", err))
	}

	selfContID := getSelfContainerID()

	if selfContID != "" {
		log.Printf("Self config id: %s", selfContID)

		env.CurrentContainerID = selfContID
	}

	log.Printf("Environment loaded: %v", env)

	return &env
}

func getSelfContainerID() string {
	hostNameCmd := exec.Command("cat", "/etc/hostname")

	selfContID, err := hostNameCmd.Output()

	if err != nil {
		log.Printf("Failed to cat /etc/hostname: %s", err)
	}

	return strings.TrimSpace(string(selfContID))
}
