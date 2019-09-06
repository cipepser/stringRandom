package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT   = "INT"
	STRING = "STRING"
	DIGIT = "d"

	PLUS     = "+"
	ASTERISK = "*"

	LBRACE = "{"
	RBRACE = "}"
	COMMA  = ","
)
