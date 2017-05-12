package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func getFileContent(fileName string) []byte {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "base64encdec: File '%v' does not exist\n", fileName)
		os.Exit(ERROR_FILE_NOT_EXIST)
	}

	// Works only for *nix systems
	fileMode := fileInfo.Mode().String()
	if len(fileMode) > 8 {
		if fileMode[1:2] != "r" && fileMode[4:5] != "r" && fileMode[7:8] != "r" {
			fmt.Fprintf(os.Stderr, "base64encdec: File '%v' is not readable\n", fileName)
			os.Exit(ERROR_FILE_READ_PERMISSION)
		}
	}

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "base64encdec: %v\n", err)
		os.Exit(ERROR_UNKNOWN)
	}

	return fileContent
}
