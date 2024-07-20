package parser

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)

type Expr interface {
	String() string
}

type Binary struct {
	left     Expr
	right    Expr
	operator token.Token
}

func NewBinary(left Expr, operator token.Token, right Expr) *Binary {
	return &Binary{
		left,
		right,
		operator,
	}
}

func (b *Binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.operator.Lexeme, b.left.String(), b.right.String())
}

type Unary struct {
	right    Expr
	operator token.Token
}

func (u *Unary) String() string {
	return fmt.Sprintf("(%s %s)", u.operator.Lexeme, u.right.String())
}

type Literal struct {
	Value string
}

func NewLiteral(literal string) *Literal {
	return &Literal{
		Value: literal,
	}
}

func (l *Literal) String() string {
	return l.Value
}

type Grouping struct {
	Value Expr
}

func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		Value: expr,
	}
}

// TODO
func (g *Grouping) String() string {
	return fmt.Sprintf("group %s", g.Value.String())
}
