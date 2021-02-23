package on_close

import (
	"fmt"
	"github.com/approvers/go-vcnotify/config"
	"os"
)

var targetFiles = config.SecretFiles


func PreClosingProcess() {
	err := deleteSecretFiles()
	if err != nil {
		fmt.Printf("Error: Unable to delete secret files. Details: %s", err)
	}
}

func deleteSecretFiles() error {
	var err error

	for _, fileName := range targetFiles {
		err = os.Remove(fileName)

		if err != nil {
			return nil
		}
	}

	return nil
}
