// Package fileinfo provides functionality for getting file information.
package fileinfo

import "os"

// GetAbsoluteFileNameFromWorkingDirectory returns for the given file name the
// absolute file name based on the working directory.
func GetAbsoluteFileNameFromWorkingDirectory(fileName string) (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", err
	}

	absoluteFileName := path + "\\" + fileName

	return absoluteFileName, nil
}
