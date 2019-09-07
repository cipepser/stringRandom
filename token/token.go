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
	// TODO: 実装する NOTDIGIT = "D"
	WORD     = "w"
	// TODO: 実装する NOTWORD  = "W"
	SPACE    = "SPACE"
	// TODO: 実装する NOTSPACE = "NOTSPACE"
	NEWLINE = "NEWLINE"
	TAB = "TAB"
	BACKSLASH = "BACKSLASH"

	// TODO: 実装する DOT = "."
	PLUS     = "+"
	ASTERISK = "*"
	// TODO: 実装する QUESTION = "?"

	LBRACE = "{"
	RBRACE = "}"
	// TODO: 実装する LPAREN = "("
	// TODO: 実装する RPAREN = ")"
	// TODO: 実装する LBRACKET = "["
	// TODO: 実装する RBRACKET = "]"
	// TODO: 実装する VERTICAL = "|"
	COMMA  = ","

	// TODO: 実装する HEAD = "^"
	// TODO: 実装する TAIL = "$"
	// TODO: 実装する？
	//   WORDHEAD = "<"
	//   WORDTAIL = ">"
	//   WORDHEADORTAIL = "b"
	//   NOTWORDHEADORTAIL = "B"
	//   \A, \z、\G
)
