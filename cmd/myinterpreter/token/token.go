package token

import "fmt"

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	//single char tokens
	LPAREN    = "LEFT_PAREN"
	RPAREN    = "RIGHT_PAREN"
	LBRACE    = "LEFT_BRACE"
	RBRACE    = "RIGHT_BRACE"
	COMMA     = "COMMA"
	DOT       = "DOT"
	MINUS     = "MINUS"
	PLUS      = "PLUS"
	SEMICOLON = "SEMICOLON"
	SLASH     = "SLASH"
	STAR      = "STAR"

	//operator
	BANG          = "BANG"
	BANG_EQUAL    = "BANG_EQUAL"
	EQUAL         = "EQUAL"
	EQUAL_EQUAL   = "EQUAL_EQUAL"
	GREATER       = "GREATER"
	GREATER_EQUAL = "GREATER_EQUAL"
	LESS          = "LESS"
	LESS_EQUAL    = "LESS_EQUAL"

	//literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	//keywords
	AND    = "AND"
	CLASS  = "CLASS"
	ELSE   = "ELSE"
	FALSE  = "FALSE"
	FUNC   = "FUN"
	FOR    = "FOR"
	IF     = "IF"
	NIL    = "NIL"
	OR     = "OR"
	PRINT  = "PRINT"
	RETURN = "RETURN"
	SUPER  = "SUPER"
	THIS   = "THIS"
	TRUE   = "TRUE"
	VAR    = "VAR"
	WHILE  = "WHILE"
)

type Token struct {
	Type      TokenType
	Lexeme    string
	Literal   interface{} //declared as object in the robert nystrom book
	Line      int
	HasError  bool
	IsComment bool
}

func New(tokType TokenType, lexeme string, literal any, line int, args ...bool) Token {
	hasError := false
	isComment := false
	if len(args) == 1 {
		hasError = args[0]
	} else if len(args) == 2 {
		isComment = args[1]
	}

	return Token{
		Type:      tokType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
		HasError:  hasError,
		IsComment: isComment,
	}
}

func (t *Token) ToString() string {
	if t.Literal == nil {
		return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, "null")
	} else {
		return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
	}
}
