package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
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

	switch command {
	case "tokenize":
		filename := os.Args[2]
		fileContents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}

		if len(fileContents) > 0 {
			s := scanner.New(string(fileContents))
			tokens, erroredTokens := s.Collect()
			s.Print(tokens)
			if len(erroredTokens) > 0 {
				os.Exit(65)
			}
		} else {
			tok := token.New(token.EOF, "", nil, 0)
			fmt.Println(tok.ToString())
		}

	case "parse":
		var toks []token.Token
		filename := os.Args[2]
		fileContents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}

		if len(fileContents) > 0 {
			s := scanner.New(string(fileContents))
			tokens, erroredTokens := s.Collect()
			toks = tokens
			//s.Print(tokens)
			if len(erroredTokens) > 0 {
				os.Exit(65)
			}
		} else {
			tok := token.New(token.EOF, "", nil, 0)
			toks = []token.Token{tok}
		}

		p := parser.New(toks)
		if err := p.Parse(); err != nil {
			fmt.Println(err.Error())
			os.Exit(65)
		}

	}
}
