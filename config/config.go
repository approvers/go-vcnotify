package config

import (
	"os"
)

const (
	TokenEnvironmentName string = "GOVCNOTIFY_DISCORD_TOKEN"
	FirebaseCredentialB64EnvironmentName string = "GOVCNOTIFY_FIREBASE_CREDENTIAL_BASE64"
	RealtimeDBURLEnvironmentName string = "GOVCNOTIFY_FIREBASE_DB_URL"

	FirebaseCredentialFileName string = "secret.json"
)

var (
	DiscordToken = os.Getenv(TokenEnvironmentName)
	FirebaseCredentialB64 = os.Getenv(FirebaseCredentialB64EnvironmentName)
	RealtimeDBURL = os.Getenv(RealtimeDBURLEnvironmentName)
)
