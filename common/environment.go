package common

import (
	"os"
	"strings"
)

func IsProduction() bool {
	serverSoftwareEnv := os.Getenv("SERVER_SOFTWARE")
	return strings.HasPrefix(serverSoftwareEnv, "Google App Engine/")
}
