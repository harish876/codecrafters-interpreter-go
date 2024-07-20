package parser

import "fmt"

type AstPrinter struct {
	expr Expr
}

func NewAstPrinter(expr Expr) *AstPrinter {
	return &AstPrinter{
		expr,
	}
}

func (p *AstPrinter) PrettyPrint() {
	fmt.Println(p.expr.String())
}
