package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/approvers/go-vcnotify/config"
)

func GetCurrentTimeOfJST() time.Time {
	nowUTC := time.Now().UTC()

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)

	return nowJST
}

func GetToken() string {
	token := os.Getenv(config.TokenEnvironmentName)

	if token == "" {
		displayError := fmt.Sprintf("Error: Unable to get environment veriable: '%s'\n",
			config.TokenEnvironmentName)

		panic(displayError)
	}

	return token
}
