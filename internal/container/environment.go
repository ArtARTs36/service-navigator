package container

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	User string `default:"developer"`
}

func InitEnvironment() *Environment {
	var env Environment

	log.Print("Loading environment")

	err := envconfig.Process("", &env)

	if err != nil {
		panic(fmt.Sprintf("Failed to load environment: %s", err))
	}

	log.Printf("Environment loaded: %v", env)

	return &env
}
