package utils

import (
	"log"
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
		log.Panicf("Error: Unable to get environment veriable: '%s'\n",
			config.TokenEnvironmentName)
	}

	return token
}
