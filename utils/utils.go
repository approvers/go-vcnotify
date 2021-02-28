package utils

import (
	b64 "encoding/base64"
	"io/ioutil"
)


func CreateFileFromB64(targetText string, outputFileName string) error {
	decodedCredentialBytes, err := b64.StdEncoding.DecodeString(targetText)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFileName, decodedCredentialBytes, 0666)
	if err != nil {
		return err
	}

	return nil
}
