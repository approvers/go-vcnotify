package config

import (
	"log"
	"os"
)

const (
	DiscordTokenEnvironmentName          string = "GOVCNOTIFY_DISCORD_TOKEN"
	FirebaseCredentialB64EnvironmentName string = "GOVCNOTIFY_FIREBASE_CREDENTIAL_BASE64"
	RealtimeDBURLEnvironmentName         string = "GOVCNOTIFY_FIREBASE_DB_URL"

	FirebaseCredentialFileName string = "secret.json"
)

var (
	SecretFiles = []string{
		FirebaseCredentialFileName,
	}
)


func GetEnvironmentVariable(variableName string) string {
	result := os.Getenv(variableName)

	if result == "" {
		log.Panicf("Error: Unable to get environment variable '%s'", variableName)
	}

	return result
}
