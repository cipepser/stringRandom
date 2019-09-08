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
	// TODO: 実装する BAR = "-" // []で使われると範囲指定子になる
	// TODO: 実装する VERTICAL = "|"
	COMMA = ","

	// TODO: 実装する HEAD = "^" // []で使われると否定子になる
	// TODO: 実装する TAIL = "$"
	// TODO: 以下は実装する？
	//   WORDHEAD = "<"
	//   WORDTAIL = ">"
	//   WORDHEADORTAIL = "b"
	//   NOTWORDHEADORTAIL = "B"
	//   \A, \z、\G
)
