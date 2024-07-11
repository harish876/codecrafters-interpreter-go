package scanner

import (
	"fmt"
	"os"
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
		line:  1,
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

func (s *Scanner) peakChar(position int) byte {
	if position >= len(s.input) {
		return '0' //EOF
	} else {
		return s.input[position]
	}
}

func (s *Scanner) NextToken() token.Token {

	var tok token.Token

	s.skipWhitespace()
	//keep a count of '\n' in the file content to track the line
	switch s.ch {
	case '(':
		tok = token.New(token.LPAREN, string(s.ch), nil, s.line)
	case ')':
		tok = token.New(token.RPAREN, string(s.ch), nil, s.line)
	case '{':
		tok = token.New(token.LBRACE, string(s.ch), nil, s.line)
	case '}':
		tok = token.New(token.RBRACE, string(s.ch), nil, s.line)
	case '=':
		if s.peakChar(s.position+1) == '=' {
			s.readChar()
			tok = token.New(token.EQUAL_EQUAL, string("=="), nil, s.line)
		} else {
			tok = token.New(token.EQUAL, string(s.ch), nil, s.line)
		}
	case '!':
		if s.peakChar(s.position+1) == '=' {
			s.readChar()
			tok = token.New(token.BANG_EQUAL, string("!="), nil, s.line)
		} else {
			tok = token.New(token.BANG, string(s.ch), nil, s.line)
		}
	case '>':
		if s.peakChar(s.position+1) == '=' {
			s.readChar()
			tok = token.New(token.GREATER_EQUAL, string(">="), nil, s.line)
		} else {
			tok = token.New(token.GREATER, string(s.ch), nil, s.line)
		}
	case '<':
		if s.peakChar(s.position+1) == '=' {
			s.readChar()
			tok = token.New(token.LESS_EQUAL, string("<="), nil, s.line)
		} else {
			tok = token.New(token.LESS, string(s.ch), nil, s.line)
		}

	case '/':
		if s.peakChar(s.position+1) == '/' {
			for range 2 {
				s.readChar()
			}
			comment := s.readComment()
			tok = token.New(
				token.IDENTFIER,
				"//"+comment,
				comment,
				s.line,
				false,
				true,
			) //temp band aid fix
		} else {
			tok = token.New(token.SLASH, string(s.ch), nil, s.line)
		}
	case ';':
		tok = token.New(token.SEMICOLON, string(s.ch), nil, s.line)
	case '+':
		tok = token.New(token.PLUS, string(s.ch), nil, s.line)
	case ',':
		tok = token.New(token.COMMA, string(s.ch), nil, s.line)
	case '*':
		tok = token.New(token.STAR, string(s.ch), nil, s.line)
	case '0':
		tok = token.New(token.EOF, "", nil, s.line)
	case '.':
		tok = token.New(token.DOT, string(s.ch), nil, s.line)
	case '-':
		tok = token.New(token.MINUS, string(s.ch), nil, s.line)
	default:
		if isLetter(s.ch) {
			identifier := s.readIdentifier()
			tok = s.fromSymbol(identifier)
		} else if s.ch == '"' {
			str := s.readString()
			tok = s.fromSymbol(str)
		} else {
			tok = token.New(token.ILLEGAL, string(s.ch), nil, s.line, true) //band-aid solution idk how to handle this
			message := fmt.Sprintf("Unexpected character: %s", string(s.ch))
			s.logError(message)
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

func (s *Scanner) readString() string {
	position := s.position
	s.readChar()
	for isLetter(s.ch) {
		s.readChar()
	}
	if s.ch != '"' {
		s.logError(fmt.Sprintf("Incorrectly terminated string: %s", s.input[position:s.position]))
	}
	return s.input[position:s.position] + string(s.ch)
}

func (s *Scanner) readComment() string {
	position := s.position
	for s.ch != '0' && s.ch != '\n' {
		s.readChar()
	}
	return s.input[position:s.position]
}

func (s *Scanner) fromSymbol(literal string) token.Token {
	lexeme := strings.ToLower(literal)
	var tok token.Token
	switch lexeme {
	case "var":
		tok = token.New(token.VAR, lexeme, nil, 1)
	default:
		if lexeme[0] == '"' && lexeme[len(lexeme)-1] == '"' {
			escapedLiteral := lexeme[1 : len(lexeme)-1]
			tok = token.New(token.STRING, lexeme, escapedLiteral, s.line)
		} else {
			tok = token.New(token.IDENTFIER, lexeme, nil, s.line)
		}
	}
	return tok
}

// have error enums
func (s *Scanner) logError(message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", s.line, message)
}

// collect all the tokens within a slice passed in
func (s *Scanner) Collect() ([]token.Token, []token.Token) {
	var validTokens []token.Token
	var erroredTokens []token.Token
	for {
		tok := s.NextToken()
		if !tok.HasError && !tok.IsComment {
			validTokens = append(validTokens, tok)
		} else if tok.HasError {
			erroredTokens = append(erroredTokens, tok)
		}

		if tok.Type == token.EOF {
			break
		}
	}
	return validTokens, erroredTokens
}

func (s *Scanner) Print(tokens []token.Token) string {
	var result string
	for _, token := range tokens {
		result += token.ToString() + "\n"
	}
	fmt.Println(result)
	return result
}
