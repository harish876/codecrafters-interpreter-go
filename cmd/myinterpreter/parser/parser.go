package parser

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/utils"
)

type Parser struct {
	tokens  []token.Token
	current int
}

func New(tokens []token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) next() token.Token {
	if p.current > len(p.tokens) {
		return p.previous()
	}
	currToken := p.tokens[p.current]
	p.current += 1
	return currToken
}

func (p *Parser) consume(tokenType token.TokenType, message string) (token.Token, error) {
	if p.check(tokenType) {
		return p.next(), nil
	} else {
		return token.New(token.ILLEGAL, "", nil, 0, true), fmt.Errorf("%s", message)
		//fmt.Errorf("token - %v, message  - %s", p.peek(), message)
	}
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) isEOF() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isEOF() {
		return false
	} else {
		return p.peek().Type == tokenType
	}
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, typ := range types {
		if p.check(typ) {
			p.next()
			return true
		}
	}
	return false
}

func (p *Parser) expression() (Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparision()
	if err != nil {
		return nil, err
	}

	for p.match(token.EQUAL_EQUAL, token.BANG_EQUAL) {
		operator := p.previous()
		right, err := p.comparision()

		if err != nil {
			return nil, err
		}

		expr = &Binary{
			left:     expr,
			operator: operator,
			right:    right,
		}

	}
	return expr, nil
}

func (p *Parser) comparision() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = &Binary{
			left:     expr,
			operator: operator,
			right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, err := p.factor()

		if err != nil {
			return nil, err
		}

		expr = &Binary{
			left:     expr,
			operator: operator,
			right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()

	if err != nil {
		return nil, err
	}

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, err := p.unary()

		if err != nil {
			return nil, err
		}
		expr = &Binary{
			left:     expr,
			operator: operator,
			right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, err := p.unary()

		if err != nil {
			return nil, err
		}

		return &Unary{
			right:    right,
			operator: operator,
		}, nil
	}
	primary, err := p.primary()
	if err != nil {
		return nil, err
	}
	return primary, nil
}

func (p *Parser) primary() (Expr, error) {
	if p.match(token.FALSE) {
		return NewLiteral("false"), nil
	} else if p.match(token.TRUE) {
		return NewLiteral("true"), nil
	} else if p.match(token.NIL) {
		return NewLiteral("nil"), nil
	}

	if p.match(token.STRING, token.NUMBER) {
		return NewLiteral(p.previous().Literal.(string)), nil
	}

	if p.match(token.LPAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(token.RPAREN, "")
		if err != nil {
			return nil, err
		}
		return NewGrouping(expr), nil
	}

	return nil, fmt.Errorf("%s", "")

}

func ReportParseError(line int, tok token.Token, message string) {
	if tok.Type == token.EOF {
		fmt.Fprintln(os.Stderr, utils.Report(line, "at end"+message))
	} else {
		fmt.Fprintln(os.Stderr, utils.Report(line, "at"+tok.Lexeme+"'"+message))
	}
}

func (p *Parser) Parse() error {
	expr, err := p.expression()
	if err != nil {
		return err
	}
	NewAstPrinter(expr).PrettyPrint()
	return nil
}
