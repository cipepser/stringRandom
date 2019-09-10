package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT       = "INT"
	STRING    = "STRING"
	DIGIT     = "d"
	NOTDIGIT  = "D"
	WORD      = "w"
	NOTWORD   = "W"
	SPACE     = "s"
	NOTSPACE  = "S"
	NEWLINE   = "n"
	TAB       = "t"
	BACKSLASH = "\\"
	// TODO: メタ文字のエスケープ処理

	DOT      = "."
	PLUS     = "+"
	ASTERISK = "*"
	// TODO: 実装する QUESTION = "?"

	LBRACE   = "{"
	RBRACE   = "}"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
	BAR      = "-" // []で使われると範囲指定子になる
	// TODO: 実装する VERTICAL = "|"
	COMMA = ","
	// TODO: 実装する NOT = "^" // []で使われると否定子になる
)
