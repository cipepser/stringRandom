package parser

import (
	"github.com/cipepser/stringRandom/ast"
	"github.com/cipepser/stringRandom/lexer"
	"testing"
)

// TODO: 複数にまたがった正規表現のテストを書く

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

func TestWordExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\w{3}`, 3, 3},
		{`\w{2,5}`, 2, 5},
		{`\w{12}`, 12, 12},
		{`\w{1,23}`, 1, 23},
		{`\w+`, 1, ast.INFINITE},
		{`\w*`, 0, ast.INFINITE},
		{`\w`, 1, 1},
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

		wordExpression, ok := stmt.Expression.(*ast.WordExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.WordExpression. got=%T]", i, stmt.Expression)
		}
		if wordExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, wordExpression.Range.Min, tt.expectedMin)
		}
		if wordExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, wordExpression.Range.Max, tt.expectedMax)
		}
	}
}

func TestSpaceExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\s{3}`, 3, 3},
		{`\s{2,5}`, 2, 5},
		{`\s{12}`, 12, 12},
		{`\s{1,23}`, 1, 23},
		{`\s+`, 1, ast.INFINITE},
		{`\s*`, 0, ast.INFINITE},
		{`\s`, 1, 1},
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

		spaceExpression, ok := stmt.Expression.(*ast.SpaceExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.SpaceExpression. got=%T]", i, stmt.Expression)
		}
		if spaceExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, spaceExpression.Range.Min, tt.expectedMin)
		}
		if spaceExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, spaceExpression.Range.Max, tt.expectedMax)
		}
	}
}

func TestNewlineExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\n{3}`, 3, 3},
		{`\n{2,5}`, 2, 5},
		{`\n{12}`, 12, 12},
		{`\n{1,23}`, 1, 23},
		{`\n+`, 1, ast.INFINITE},
		{`\n*`, 0, ast.INFINITE},
		{`\n`, 1, 1},
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

		newlineExpression, ok := stmt.Expression.(*ast.NewlineExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.NewlineExpression. got=%T]", i, stmt.Expression)
		}
		if newlineExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, newlineExpression.Range.Min, tt.expectedMin)
		}
		if newlineExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, newlineExpression.Range.Max, tt.expectedMax)
		}
	}
}
