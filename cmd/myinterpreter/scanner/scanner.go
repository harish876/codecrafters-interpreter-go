package scanner

import (
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)

//scanner is called the lexer or the parser

type Scanner struct {
	input        string
	position     int  //current position in the input - points to current char
	readPosition int  //current reading position - after the current char
	ch           byte //current character under position
	line         int
}

func New(input string) *Scanner {
	s := &Scanner{
		input: input,
	}
	s.readChar()
	return s
}

// this corresponds to the advance function in the robert nystrom book
func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input) {
		s.ch = '0'
	} else {
		s.ch = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition += 1
}

func (s *Scanner) NextToken() token.Token {

	var tok token.Token

	s.skipWhitespace()
	switch s.ch {
	case '(':
		tok = token.New(token.LPAREN, string(s.ch), nil, 1)
	case ')':
		tok = token.New(token.RPAREN, string(s.ch), nil, 1)
	case '{':
		tok = token.New(token.LBRACE, string(s.ch), nil, 1)
	case '}':
		tok = token.New(token.RBRACE, string(s.ch), nil, 1)
	case '=':
		tok = token.New(token.EQUAL, string(s.ch), nil, 1)
	case ';':
		tok = token.New(token.SEMICOLON, string(s.ch), nil, 1)
	case '+':
		tok = token.New(token.PLUS, string(s.ch), nil, 1)
	case ',':
		tok = token.New(token.COMMA, string(s.ch), nil, 1)
	case '*':
		tok = token.New(token.STAR, string(s.ch), nil, 1)
	case '0':
		tok = token.New(token.EOF, "", nil, 1)
	case '.':
		tok = token.New(token.DOT, string(s.ch), nil, 1)
	case '-':
		tok = token.New(token.MINUS, string(s.ch), nil, 1)
	default:
		if isLetter(s.ch) {
			tok = fromSymbol(s.readIdentifier())
		} else {
			tok = token.New(token.ILLEGAL, string(s.ch), nil, 1)
		}
	}
	s.readChar()
	return tok
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\r' || s.ch == '\n' {
		s.readChar()
	}
}

func (s *Scanner) readIdentifier() string {
	position := s.position
	for isLetter(s.ch) {
		s.readChar()
	}
	return s.input[position:s.position]

}

func fromSymbol(literal string) token.Token {
	lexeme := strings.ToLower(literal)
	var tok token.Token
	switch lexeme {
	case "var":
		tok = token.New(token.VAR, lexeme, nil, 1)
	default:
		if len(lexeme) >= 1 && lexeme[0] == '"' && lexeme[len(lexeme)-1] == '"' {
			escapedLiteral := strings.ReplaceAll(literal, `"`, "")
			tok = token.New(token.STRING, escapedLiteral, escapedLiteral, 1)
		} else {
			tok = token.New(token.IDENTFIER, lexeme, nil, 1)
		}
	}
	return tok
}
