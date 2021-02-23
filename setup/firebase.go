package setup

import (
	"context"
	"log"

	"github.com/approvers/go-vcnotify/config"
	"github.com/approvers/go-vcnotify/utils"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var (
	ctx = context.Background()
	fileName = config.FirebaseCredentialFileName
	credentialBase64 = config.GetEnvironmentVariable(config.FirebaseCredentialB64EnvironmentName)
	DatabaseURL = config.GetEnvironmentVariable(config.RealtimeDBURLEnvironmentName)
)


func GetDBClient() *db.Client {
	err := createCredentialFile()
	if err != nil {
		log.Panicf("Error: An error has occured while creating a credential from base64.\n" +
			"Detail:%s\n", err)
	}

	opt := getOptionWithCredential()
	app := getFirebaseApp(opt)
	client := _getDBClient(app)

	return client
}


func createCredentialFile() error {
	err := utils.CreateFileFromB64(credentialBase64, fileName)

	if err != nil {
		return err
	}

	return nil
}


func getOptionWithCredential() option.ClientOption {
	opt := option.WithCredentialsFile(fileName)

	return opt
}


func getFirebaseApp(opt option.ClientOption) *firebase.App {
	conf := &firebase.Config {
		DatabaseURL: DatabaseURL,
	}

	app, err := firebase.NewApp(ctx, conf, opt)

	if err != nil {
		log.Panicf("Error: An error has occured while initializing a firebase app.\n" +
			"Detail:%s\n", err)
	}

	return app
}


func _getDBClient(app *firebase.App) *db.Client {
	client, err := app.Database(ctx)

	if err != nil {
		log.Panicf("Error: An error has occured while creating a database client.\n" +
			"Detail:%s\n", err)
	}

	return client
}
