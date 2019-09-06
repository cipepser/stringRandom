package lexer

import (
	"github.com/cipepser/stringRandom/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `\d{3}\d{2,5}\d{12}\d{1,23}\d+\d*Hogee*\d\d\w`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DIGIT, "d"},
		{token.LBRACE, "{"},
		{token.INT, "3"},
		{token.RBRACE, "}"},

		{token.DIGIT, "d"},
		{token.LBRACE, "{"},
		{token.INT, "2"},
		{token.COMMA, ","},
		{token.INT, "5"},
		{token.RBRACE, "}"},

		{token.DIGIT, "d"},
		{token.LBRACE, "{"},
		{token.INT, "12"},
		{token.RBRACE, "}"},

		{token.DIGIT, "d"},
		{token.LBRACE, "{"},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "23"},
		{token.RBRACE, "}"},

		{token.DIGIT, "d"},
		{token.PLUS, "+"},

		{token.DIGIT, "d"},
		{token.ASTERISK, "*"},

		{token.STRING, "Hoge"},

		{token.STRING, "e"},
		{token.ASTERISK, "*"},

		{token.DIGIT, "d"},
		{token.DIGIT, "d"},

		{token.WORD, "w"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d - tokentype wrong. expected=%q, got=%q]", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d - literal wrong. expected=%q, got=%q]", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
