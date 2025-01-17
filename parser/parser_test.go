package parser

import (
	"testing"

	"github.com/piyushgupta53/go-monkey/ast"
	"github.com/piyushgupta53/go-monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		t.Errorf("parser had %d errors:", len(p.Errors()))
		for _, err := range p.Errors() {
			t.Errorf("parser error: %q", err)
		}
		t.FailNow()
	}

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d\nstatements=%v\n",
			len(program.Statements), program.Statements)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestLetStatementErrors(t *testing.T) {
	input := `
let x 5;
let = 10;
let 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if len(p.Errors()) == 0 {
		t.Error("parser didn't report any errors for invalid input")
	}

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
}
