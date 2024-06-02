package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port         string
	RunningOnK8S bool
	Env          map[string]string
}

func NewConfig() (Config, error) {
	config := Config{
		Env: make(map[string]string),
	}

	config.Port = os.Getenv("ECHO_SERVER_LISTEN_PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	if _, err := strconv.Atoi(config.Port); err != nil {
		return config, fmt.Errorf("`ECHO_SERVER_LISTEN_PORT` must be an integer. Received %q", config.Port)
	}

	for _, envString := range os.Environ() {
		if !strings.HasPrefix(envString, "KUBERNETES_") {
			continue
		}

		splits := strings.SplitN(envString, "=", 2)

		config.Env[splits[0]] = splits[1]
	}

	_, config.RunningOnK8S = config.Env["KUBERNETES_SERVICE_HOST"]

	return config, nil
}
