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
	SLASH     = "SLAS"
	STAR      = "STAR"

	//operator
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	//literals
	IDENTFIER = "IDENTFIER"
	STRING    = "STRING"
	NUMBER    = "NUMBER"

	//keywords
	AND    = "AND"
	CLASS  = "CLASS"
	ELSE   = "ELSE"
	FALSE  = "FALSE"
	FUNC   = "FUNC"
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
	Type     TokenType
	Lexeme   string
	Literal  interface{} //declared as object in the robert nystrom book
	Line     int
	HasError bool
}

func New(tokType TokenType, lexeme string, literal any, line int, args ...any) Token {
	hasError := false
	if len(args) > 0 {
		hasError = args[0].(bool)
	}

	return Token{
		Type:     tokType,
		Lexeme:   lexeme,
		Literal:  literal,
		Line:     line,
		HasError: hasError,
	}
}

func (t *Token) ToString() string {
	if t.Literal == nil {
		return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, "null")
	} else {
		return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
	}
}
