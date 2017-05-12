package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func encode(fileContent []byte) string {
	return base64.StdEncoding.EncodeToString(fileContent)
}

func decode(fileContent []byte) string {
	fileContentString := string(fileContent[:])
	data, err := base64.StdEncoding.DecodeString(fileContentString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "base64encdec: %v\n", err)
		os.Exit(ERROR_UNKNOWN)
	}

	return string(data[:])
}
