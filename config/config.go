package config


const (
	DiscordTokenEnvironmentName          string = "GOVCNOTIFY_DISCORD_TOKEN"
	FirebaseCredentialB64EnvironmentName string = "GOVCNOTIFY_FIREBASE_CREDENTIAL_BASE64"
	RealtimeDBURLEnvironmentName         string = "GOVCNOTIFY_FIREBASE_DB_URL"

	FirebaseCredentialFileName string = "secret.json"

	RealtimeDBGuildDirectoryName string = "guilds"
)

var (
	NamesOfSecretFiles = []string{
		FirebaseCredentialFileName,
	}
)
