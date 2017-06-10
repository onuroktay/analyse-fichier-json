package OnurTPIjsonReader

import "os"

// checkIfFileExist verify if a file exists on disk
func checkIfFileExist(filename string) (fileExists bool) {
	if _, err := os.Stat(filename); err == nil {
		fileExists = true
	}

	return
}
