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
