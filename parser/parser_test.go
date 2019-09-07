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

func TestNotDigitExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\D{3}`, 3, 3},
		{`\D{2,5}`, 2, 5},
		{`\D{12}`, 12, 12},
		{`\D{1,23}`, 1, 23},
		{`\D+`, 1, ast.INFINITE},
		{`\D*`, 0, ast.INFINITE},
		{`\D`, 1, 1},
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

		notDigitExpression, ok := stmt.Expression.(*ast.NotDigitExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.NotDigtExpression. got=%T]", i, stmt.Expression)
		}
		if notDigitExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, notDigitExpression.Range.Min, tt.expectedMin)
		}
		if notDigitExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, notDigitExpression.Range.Max, tt.expectedMax)
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

func TestNotWordExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\W{3}`, 3, 3},
		{`\W{2,5}`, 2, 5},
		{`\W{12}`, 12, 12},
		{`\W{1,23}`, 1, 23},
		{`\W+`, 1, ast.INFINITE},
		{`\W*`, 0, ast.INFINITE},
		{`\W`, 1, 1},
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

		notWordExpression, ok := stmt.Expression.(*ast.NotWordExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.NotWordExpression. got=%T]", i, stmt.Expression)
		}
		if notWordExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, notWordExpression.Range.Min, tt.expectedMin)
		}
		if notWordExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, notWordExpression.Range.Max, tt.expectedMax)
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

func TestTabExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\t{3}`, 3, 3},
		{`\t{2,5}`, 2, 5},
		{`\t{12}`, 12, 12},
		{`\t{1,23}`, 1, 23},
		{`\t+`, 1, ast.INFINITE},
		{`\t*`, 0, ast.INFINITE},
		{`\t`, 1, 1},
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

		tabExpression, ok := stmt.Expression.(*ast.TabExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.TabExpression. got=%T]", i, stmt.Expression)
		}
		if tabExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, tabExpression.Range.Min, tt.expectedMin)
		}
		if tabExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, tabExpression.Range.Max, tt.expectedMax)
		}
	}
}

func TestBackslashExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`\\{3}`, 3, 3},
		{`\\{2,5}`, 2, 5},
		{`\\{12}`, 12, 12},
		{`\\{1,23}`, 1, 23},
		{`\\+`, 1, ast.INFINITE},
		{`\\*`, 0, ast.INFINITE},
		{`\\`, 1, 1},
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

		backslashExpression, ok := stmt.Expression.(*ast.BackslashExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.BackslashExpression. got=%T]", i, stmt.Expression)
		}
		if backslashExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, backslashExpression.Range.Min, tt.expectedMin)
		}
		if backslashExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, backslashExpression.Range.Max, tt.expectedMax)
		}
	}
}

func TestDotExpression(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{`.{3}`, 3, 3},
		{`.{2,5}`, 2, 5},
		{`.{12}`, 12, 12},
		{`.{1,23}`, 1, 23},
		{`.+`, 1, ast.INFINITE},
		{`.*`, 0, ast.INFINITE},
		{`.`, 1, 1},
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

		dotExpression, ok := stmt.Expression.(*ast.DotExpression)
		if !ok {
			t.Fatalf("test[%d - exp not *ast.DotExpression. got=%T]", i, stmt.Expression)
		}
		if dotExpression.Range.Min != tt.expectedMin {
			t.Fatalf("test[%d - wrong Range(min). got=%v, want=%v]", i, dotExpression.Range.Min, tt.expectedMin)
		}
		if dotExpression.Range.Max != tt.expectedMax {
			t.Fatalf("test[%d - wrong Range(max). got=%v, want=%v]", i, dotExpression.Range.Max, tt.expectedMax)
		}
	}
}
