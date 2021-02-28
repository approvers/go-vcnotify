package config

import (
	"log"
	"os"
)


func GetEnvironmentVariable(variableName string) string {
	result := os.Getenv(variableName)

	if result == "" {
		log.Panicf("Error: Unable to get environment variable '%s'", variableName)
	}

	return result
}
