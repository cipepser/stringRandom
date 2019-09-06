package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT      = "INT"
	STRING   = "STRING"
	DIGIT    = "d"
	//NOTDIGIT = "D"
	WORD     = "w"
	//NOTWORD  = "W"

	PLUS     = "+"
	ASTERISK = "*"

	LBRACE = "{"
	RBRACE = "}"
	COMMA  = ","
)
