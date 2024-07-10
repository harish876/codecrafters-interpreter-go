package scanner

import (
	"fmt"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	s := New(input)

	expectedResult := []token.Token{
		{
			Type:    token.EQUAL,
			Lexeme:  "=",
			Literal: nil,
		},
		{
			Type:    token.PLUS,
			Lexeme:  "+",
			Literal: nil,
		},
		{
			Type:    token.LPAREN,
			Lexeme:  "(",
			Literal: nil,
		},
		{
			Type:    token.RPAREN,
			Lexeme:  ")",
			Literal: nil,
		},
		{
			Type:    token.LBRACE,
			Lexeme:  "{",
			Literal: nil, //this is weird i dont understand
		},
		{
			Type:    token.RBRACE,
			Lexeme:  "}",
			Literal: nil, //this is weird i dont understand
		},
		{
			Type:    token.COMMA,
			Lexeme:  ",",
			Literal: nil,
		},
		{
			Type:    token.SEMICOLON,
			Lexeme:  ";",
			Literal: nil,
		},
		{
			Type:    token.EOF,
			Lexeme:  "",
			Literal: nil,
		},
	}

	for _, result := range expectedResult {
		tok := s.NextToken()
		if result.Type != tok.Type {
			t.Fatalf("Expected token type - %v recieved token type - %v", result.Type, tok.Type)
		}

		if result.Lexeme != tok.Lexeme {
			t.Fatalf(
				"Expected token Lexeme - %v recieved token Lexeme - %s",
				result.Lexeme,
				tok.Lexeme,
			)
		}

		if result.Literal != tok.Literal {
			t.Fatalf(
				"Expected token literal - %v recieved token literal - %v",
				result.Literal,
				tok.Literal,
			)
		}
	}

	fmt.Println("Test 1 Ran Successfully.")
}

/*
func TestNextTokenV1(t *testing.T) {
	input := `var language = "lox";`
	s := New(input)
	_ = s

	expectedResult := []token.Token{
		{
			Type:    token.VAR,
			Lexeme:  "var",
			Literal: nil,
		},
		{
			Type:    token.IDENTFIER,
			Lexeme:  "language",
			Literal: nil,
		},
		{
			Type:    token.EQUAL,
			Lexeme:  "=",
			Literal: nil,
		},
		{
			Type:    token.STRING,
			Lexeme:  "lox",
			Literal: "lox", //this is weird i dont understand
		},
		{
			Type:    token.SEMICOLON,
			Lexeme:  ";",
			Literal: nil,
		},
		{
			Type:    token.EOF,
			Lexeme:  "",
			Literal: nil,
		},
	}

	tokens := make([]token.Token, 0)

	for _, result := range expectedResult {
		tok := s.NextToken()
		tokens = append(tokens, tok)
		if result.Type != tok.Type {
			t.Fatalf("Expected token type - %v recieved token type - %v\n", result.Type, tok.Type)
		}

		if result.Lexeme != tok.Lexeme {
			t.Fatalf(
				"Expected token Lexeme - %s recieved token Lexeme - %s\n",
				result.Lexeme,
				tok.Lexeme,
			)
		}

		if result.Literal != tok.Literal {
			t.Fatalf(
				"Expected token literal - %v recieved token literal - %v\n",
				result.Literal,
				tok.Literal,
			)
		}
	}
	fmt.Println(tokens)
}
*/

func TestNextTokenV2(t *testing.T) {
	input := ""
	s := New(input)
	_ = s

	expectedResult := []token.Token{
		{
			Type:    token.EOF,
			Lexeme:  "",
			Literal: nil,
		},
	}

	tokens := make([]token.Token, 0)

	for _, result := range expectedResult {
		tok := s.NextToken()
		tokens = append(tokens, tok)
		if result.Type != tok.Type {
			t.Fatalf("Expected token type - %v recieved token type - %v\n", result.Type, tok.Type)
		}

		if result.Lexeme != tok.Lexeme {
			t.Fatalf(
				"Expected token Lexeme - %s recieved token Lexeme - %s\n",
				result.Lexeme,
				tok.Lexeme,
			)
		}

		if result.Literal != tok.Literal {
			t.Fatalf(
				"Expected token literal - %v recieved token literal - %v\n",
				result.Literal,
				tok.Literal,
			)
		}
	}
	fmt.Println("Test 3 ran Successfully")
	fmt.Println(tokens[0].ToString())
}

func TestNextTokenV3(t *testing.T) {
	input := `(()`
	s := New(input)

	expectedResult := []token.Token{
		{
			Type:    token.LPAREN,
			Lexeme:  "(",
			Literal: nil,
		},
		{
			Type:    token.LPAREN,
			Lexeme:  "(",
			Literal: nil,
		}, {
			Type:    token.RPAREN,
			Lexeme:  ")",
			Literal: nil,
		},
		{
			Type:    token.EOF,
			Lexeme:  "",
			Literal: nil,
		},
	}

	tokens := make([]token.Token, 0)

	for _, result := range expectedResult {
		tok := s.NextToken()
		tokens = append(tokens, tok)
		if result.Type != tok.Type {
			t.Fatalf("Expected token type - %v recieved token type - %v", result.Type, tok.Type)
		}

		if result.Lexeme != tok.Lexeme {
			t.Fatalf(
				"Expected token Lexeme - %v recieved token Lexeme - %s",
				result.Lexeme,
				tok.Lexeme,
			)
		}

		if result.Literal != tok.Literal {
			t.Fatalf(
				"Expected token literal - %v recieved token literal - %v",
				result.Literal,
				tok.Literal,
			)
		}
	}

	fmt.Println("Test 4 Ran Successfully.")
	var result string
	for _, token := range tokens {
		result += token.ToString() + "\n"
	}
	fmt.Println(result)
}
