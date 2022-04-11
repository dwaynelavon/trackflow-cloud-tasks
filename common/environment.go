package common

import (
	"os"
)

func IsProduction() bool {
	env := os.Getenv("ENVIRONMENT")
	return env == "PROD"
}
