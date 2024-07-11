package scanner

import (
	"fmt"
	"os"

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
		s.ch = 0
	} else {
		s.ch = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition += 1
}

func (s *Scanner) peakChar() byte {
	position := s.position + 1
	if position >= len(s.input) {
		return 0 //EOF
	} else {
		return s.input[position]
	}
}

func (s *Scanner) rewind() {
	currPosition := s.position
	if currPosition > 0 {
		s.position -= 1
		s.ch = s.input[s.position]
		s.readPosition = currPosition
	} else {
		fmt.Println("Unable to rewind. No character to rewind")
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
		if s.peakChar() == '=' {
			s.readChar()
			tok = token.New(token.EQUAL_EQUAL, string("=="), nil, s.line)
		} else {
			tok = token.New(token.EQUAL, string(s.ch), nil, s.line)
		}
	case '!':
		if s.peakChar() == '=' {
			s.readChar()
			tok = token.New(token.BANG_EQUAL, string("!="), nil, s.line)
		} else {
			tok = token.New(token.BANG, string(s.ch), nil, s.line)
		}
	case '>':
		if s.peakChar() == '=' {
			s.readChar()
			tok = token.New(token.GREATER_EQUAL, string(">="), nil, s.line)
		} else {
			tok = token.New(token.GREATER, string(s.ch), nil, s.line)
		}
	case '<':
		if s.peakChar() == '=' {
			s.readChar()
			tok = token.New(token.LESS_EQUAL, string("<="), nil, s.line)
		} else {
			tok = token.New(token.LESS, string(s.ch), nil, s.line)
		}
	case '/':
		if s.peakChar() == '/' {
			comment := s.readComment()
			tok = token.New(
				token.IDENTIFIER,
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
	case 0:
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
			tok = s.readString()
		} else if isNumber(s.ch) {
			tok = s.readNumber()
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
		if s.ch == '\n' {
			s.line += 1
		}
		s.readChar()
	}
}

func (s *Scanner) readIdentifier() string {
	position := s.position
	for isLetter(s.ch) || isNumber(s.ch) {
		s.readChar()
	}
	//do i need to rewind here?
	currentPosition := s.position
	s.rewind()
	return s.input[position:currentPosition]
}

// single line string literals
func (s *Scanner) readString() token.Token {
	position := s.position
	s.readChar()
	for s.ch != '"' && s.ch != '\n' && s.ch != 0 {
		s.readChar()
	}
	if s.ch != '"' {
		s.logError("Unterminated string.")
		return token.New(token.STRING, "", nil, s.line, true) //todo update lexeme
	}
	lexeme := s.input[position:s.position] + string(s.ch)
	return token.New(token.STRING, lexeme, lexeme[1:len(lexeme)-1], s.line, false)
}

func (s *Scanner) readNumber() token.Token {
	position := s.position
	dotCount := 0
	for isNumber(s.ch) {
		s.readChar()
	}
	if s.ch == '.' && isNumber(s.peakChar()) {
		dotCount += 1
		s.readChar()
		for isNumber(s.ch) {
			s.readChar()
		}
	}
	numLexeme := s.input[position:s.position]
	s.rewind()
	literal := numLexeme

	//remove trailing zeroes
	if dotCount > 0 {
		endIdx := len(literal)
		for endIdx > 0 && literal[endIdx-1] == '0' {
			endIdx -= 1
		}
		if endIdx > 0 && literal[endIdx-1] == '.' {
			endIdx -= 1
			dotCount -= 1
		}
		literal = literal[:endIdx]
	}
	if dotCount == 0 {
		literal += ".0"
	}
	return token.New(token.NUMBER, numLexeme, literal, s.line)
}

func (s *Scanner) readComment() string {
	position := s.position
	for s.ch != 0 && s.ch != '\n' {
		s.readChar()
	}
	s.line += 1
	return s.input[position:s.position]
}

func (s *Scanner) fromSymbol(literal string) token.Token {
	lexeme := literal
	var tok token.Token
	switch lexeme {
	case "var":
		tok = token.New(token.VAR, lexeme, nil, s.line)
	case "and":
		tok = token.New(token.AND, lexeme, nil, s.line)
	case "class":
		tok = token.New(token.CLASS, lexeme, nil, s.line)
	case "else":
		tok = token.New(token.ELSE, lexeme, nil, s.line)
	case "false":
		tok = token.New(token.FALSE, lexeme, nil, s.line)
	case "true":
		tok = token.New(token.TRUE, lexeme, nil, s.line)
	case "for":
		tok = token.New(token.FOR, lexeme, nil, s.line)
	case "fun":
		tok = token.New(token.FUNC, lexeme, nil, s.line)
	case "if":
		tok = token.New(token.IF, lexeme, nil, s.line)
	case "nil":
		tok = token.New(token.NIL, lexeme, nil, s.line)
	case "or":
		tok = token.New(token.OR, lexeme, nil, s.line)
	case "print":
		tok = token.New(token.PRINT, lexeme, nil, s.line)
	case "return":
		tok = token.New(token.RETURN, lexeme, nil, s.line)
	case "this":
		tok = token.New(token.THIS, lexeme, nil, s.line)
	case "while":
		tok = token.New(token.WHILE, lexeme, nil, s.line)
	case "super":
		tok = token.New(token.SUPER, lexeme, nil, s.line)
	default:
		tok = token.New(token.IDENTIFIER, lexeme, nil, s.line)

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

func isNumber(char byte) bool {
	return (char >= '0' && char <= '9')
}
