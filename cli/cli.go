package cli

import "os"

func ConfigurationDir() string {
	return os.Getenv("DOCKER_CONFIG")
}
