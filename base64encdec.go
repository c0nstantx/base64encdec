package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get action
	if (len(os.Args[1:])) == 0 {
		fmt.Fprintf(os.Stderr, "base64encdec: You have to select an action\n")
		printHelp()
		os.Exit(ERROR_ACTION_NOT_DEFINED)
	}

	action := os.Args[1]
	if action == "help" {
		printHelp()
		os.Exit(ERROR_NO_ERROR)
	}
	if strings.Compare(action, "encode") != 0 && strings.Compare(action, "decode") != 0 {
		fmt.Fprintf(os.Stderr, "base64encdec: The selected action '%s' is not valid. Available actions are 'encode' and 'decode'\n", action)
		printHelp()
		os.Exit(ERROR_ACTION_NOT_EXIST)
	}

	// Get file
	if (len(os.Args[2:])) == 0 {
		fmt.Fprintln(os.Stderr, "base64encdec: You have to select a file to encode")
		printHelp()
		os.Exit(ERROR_FILE_NOT_DEFINED)
	}
	fileName := os.Args[2]

	var result string
	if action == "encode" {
		result = encode(getFileContent(fileName))
	} else {
		result = decode(getFileContent(fileName))
	}

	fmt.Println(result)
}

func printHelp() {
	fmt.Fprintln(os.Stdout, "WELCOME TO base64 enc/decoder\n\nUsage:\nbase64encdec [encode|decode] file\t\t\t\t\tEncode/Decode files\nbase64encdec help\t\t\t\t\t\t\tPrint this help file\n\nExamples:\nbase64encdec encode /path/to/file\t\t\t\t\tPrint to stdout\t(encode)\nbase64encdec encode /path/to/file > /path/to/encoded/file\t\tSave to file\t(encode)\nbase64encdec decode /path/to/file\t\t\t\t\tPrint to stdout\t(decode)\nbase64encdec decode /path/to/file > /path/to/decoded/file\t\tSave to file\t(decode)\n")
}
