package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT = "INT"

	DIGIT = "d"

	LBRACE = "{"
	RBRACE = "}"
	COMMA  = ","
)
