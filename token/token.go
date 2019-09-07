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

	DOT      = "."
	PLUS     = "+"
	ASTERISK = "*"
	// TODO: 実装する QUESTION = "?"

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"
	// TODO: 実装する LBRACKET = "["
	// TODO: 実装する RBRACKET = "]"
	// TODO: 実装する VERTICAL = "|"
	COMMA = ","
	// TODO: メタ文字のエスケープ処理

	// TODO: 実装する HEAD = "^" // []で使われると否定子になる
	// TODO: 実装する TAIL = "$"
	// TODO: 実装する？
	//   WORDHEAD = "<"
	//   WORDTAIL = ">"
	//   WORDHEADORTAIL = "b"
	//   NOTWORDHEADORTAIL = "B"
	//   \A, \z、\G
)
