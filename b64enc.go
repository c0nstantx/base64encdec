package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c0nstantx/base64encdec/encoder"
)

const ERROR_FILE_NOT_DEFINED = 1
const ERROR_FILE_NOT_EXIST = 2
const ERROR_FILE_READ_PERMISSION = 3
const ERROR_UNKNOWN = 4

func main() {
	if (len(os.Args[1:])) == 0 {
		fmt.Fprintln(os.Stderr, "b64enc: You have to select a file to encode")
		os.Exit(ERROR_FILE_NOT_DEFINED)
	}
	fileName := os.Args[1]
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "b64enc: File '%v' does not exist\n", fileName)
		os.Exit(ERROR_FILE_NOT_EXIST)
	}

	// Works only for *nix systems
	fileMode := fileInfo.Mode().String()
	if len(fileMode) > 8 {
		if fileMode[1:2] != "r" && fileMode[4:5] != "r" && fileMode[7:8] != "r" {
			fmt.Fprintf(os.Stderr, "b64enc: File '%v' is not readable\n", fileName)
			os.Exit(ERROR_FILE_READ_PERMISSION)
		}
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "b64enc: %v\n", err)
		os.Exit(ERROR_UNKNOWN)
	}

	encoded := encoder.Encode(data)

	fmt.Println(encoded)
}
