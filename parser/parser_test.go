package parser

import (
	"github.com/cipepser/stringRandom/ast"
	"github.com/cipepser/stringRandom/lexer"
	"testing"
)

func TestDigitExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\d{3}`, 3, 3},
		{`\d{2,5}`, 2, 5},
		{`\d{12}`, 12, 12},
		{`\d{1,23}`, 1, 23},
		{`\d+`, 1, ast.INFINITE},
		{`\d*`, 0, ast.INFINITE},
		{`\d`, 1, 1},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.Parse()

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("test[%d - program.Statement[0] is not ast.ExpressionStatement. got=%T]",
				i, program.Statements[0])
		}

		digitExpression, ok := stmt.Expression.(*ast.DigitExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.DigtExpression. got=%T]", i, stmt.Expression)
		}
		if digitExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, digitExpression.Range.Min, tt.expectedMin)
		}
		if digitExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, digitExpression.Range.Max, tt.expectedMax)
		}
	}
}

func TestStringExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedLiteral string
		expectedMin     int
		expectedMax     int
	}{
		{`Hoge`, "Hoge", 1, 1},
		{`Hogee*`, "Hoge", 1, 1},
		{`d{1,2}`, "d", 1, 2},
		{`a+`, "a", 1, ast.INFINITE},
		{`b*`, "b", 0, ast.INFINITE},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.Parse()

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("test[%d - program.Statement[0] is not ast.ExpressionStatement. got=%T]",
				i, program.Statements[0])
		}

		stringExpression, ok := stmt.Expression.(*ast.StringExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.StringExpression. got=%T]", i, stmt.Expression)
		}
		if stringExpression.TokenLiteral() != tt.expectedLiteral {
			t.Fatalf("test[%d - wrong token literal. got=%s, want=%s]", i, stringExpression.TokenLiteral(), tt.expectedLiteral)
		}
		if stringExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, stringExpression.Range.Min, tt.expectedMin)
		}
		if stringExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, stringExpression.Range.Max, tt.expectedMax)
		}
	}
}
