package exit

import (
	"fmt"
)


func PreExitProcess() (err error) {
	err = deleteSecretFiles()
	if err != nil {
		fmt.Printf("Error: Unable to delete secret files. Details: %s", err)
	}

	return err
}
