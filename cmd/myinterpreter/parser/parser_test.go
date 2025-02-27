package parser

import (
	"fmt"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

/*
func TestBinaryExpr(t *testing.T) {
	left := &Literal{Value: "4.0"}
	operator := token.New(token.PLUS, "+", nil, 0)
	right := &Literal{Value: "5.0"}

	expr := NewBinary(left, operator, right)
	NewAstPrinter(expr).PrettyPrint()
}
*/

func TestPrimary1(t *testing.T) {
	s := scanner.New(`true`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary2(t *testing.T) {
	s := scanner.New(`false`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary3(t *testing.T) {
	s := scanner.New(`2`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary4(t *testing.T) {
	s := scanner.New(`"hello world"`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary5(t *testing.T) {
	s := scanner.New(`nil`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary6(t *testing.T) {
	s := scanner.New(`2+3`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary7(t *testing.T) {
	s := scanner.New(`1*2+3`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary8(t *testing.T) {
	s := scanner.New(`("foo")`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestPrimary9(t *testing.T) {
	s := scanner.New(`("foo"`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		fmt.Println(err)
	}
}

func TestPrimary10(t *testing.T) {
	s := scanner.New(`()`)
	tokens, erroredTokens := s.Collect()
	if len(erroredTokens) > 0 {
		t.Fatalf("Code contains some errored tokens - %v", erroredTokens)
	}
	for _, tok := range tokens {
		fmt.Println(tok.Type)
	}
	p := New(tokens)
	if err := p.Parse(); err != nil {
		t.Fatalf("%v", err)
	}
}
