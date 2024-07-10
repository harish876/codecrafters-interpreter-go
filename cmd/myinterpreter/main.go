package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		s := scanner.New(string(fileContents))
		tokens := make([]token.Token, 0)
		for {
			tok := s.NextToken()
			tokens = append(tokens, tok)
			if tok.Type == token.EOF {
				break
			}
		}

		var result string
		for _, token := range tokens {
			result += token.ToString()
		}
		fmt.Println(result)
	} else {
		tok := token.New(token.EOF, "", nil, 0)
		fmt.Println(tok.ToString())
	}
}
