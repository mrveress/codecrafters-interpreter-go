package main

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/interpreter"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 2 {
		interpreter.Errorln(1, "Usage: ./your_program.sh <command - required> <filename>")
	}

	command := os.Args[1]

	if command == "tokenize" {
		if len(os.Args) < 3 {
			interpreter.Errorln(1, "Usage: ./your_program.sh tokenize <filename>")
		}

		filename := os.Args[2]
		fileContents := getFileContents(filename)

		scanner := interpreter.NewScanner(fileContents)
		scanner.ScanTokens()
		scanner.PrintLines()

		os.Exit(scanner.GetExitCode())
	} else if command == "parse" {
		if len(os.Args) < 3 {
			interpreter.Errorln(1, "Usage: ./your_program.sh parse <filename>")
		}

		filename := os.Args[2]
		fileContents := getFileContents(filename)

		scanner := interpreter.NewScanner(fileContents)
		scanner.ScanTokens()

		parser := interpreter.NewParser(scanner.Tokens)
		interpreter.AstPrinter{}.Print(parser.Parse())
	} else if command == "evaluate" {
		if len(os.Args) < 3 {
			interpreter.Errorln(1, "Usage: ./your_program.sh evaluate <filename>")
		}

		filename := os.Args[2]
		fileContents := getFileContents(filename)

		scanner := interpreter.NewScanner(fileContents)
		scanner.ScanTokens()

		parser := interpreter.NewParser(scanner.Tokens)
		expression := parser.Parse()

		interpret := interpreter.NewInterpreter()
		result := interpret.Interpret(expression)

		interpreter.Fprintf(os.Stdout, "%s\n", result)
	} else {
		interpreter.Errorf(1, "Unknown command: %s\n", command)
	}
}

func getFileContents(filename string) string {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		interpreter.Errorf(1, "Error reading file: %v\n", err)
	}
	return string(fileContents)
}
