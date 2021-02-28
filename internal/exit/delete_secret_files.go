package exit

import (
	"github.com/approvers/go-vcnotify/config"
	"os"
)

var targetFiles = config.NamesOfSecretFiles


func deleteSecretFiles() (err error) {
	for _, fileName := range targetFiles {
		err = os.Remove(fileName)

		if err != nil {
			return nil
		}
	}

	return nil
}
