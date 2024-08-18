package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/interpreter"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh <command>* <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command == "tokenize" {
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
			os.Exit(1)
		}

		filename := os.Args[2]
		fileContents := getFileContents(filename)

		scanner := interpreter.NewScanner(fileContents)
		scanner.ScanTokens()
		scanner.PrintLines()
		os.Exit(scanner.GetExitCode())
	} else if command == "parse" {
		filename := os.Args[2]
		fileContents := getFileContents(filename)

		scanner := interpreter.NewScanner(fileContents)
		scanner.ScanTokens()

		parser := interpreter.NewParser(scanner.Tokens)
		interpreter.AstPrinter{}.Print(parser.Parse())
	} else {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func getFileContents(filename string) string {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	return string(fileContents)
}
