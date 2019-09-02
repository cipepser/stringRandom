package lexer

import "github.com/cipepser/stringRandom/token"

type Lexer struct {
	b            string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{b: input}
	l.readPosition = l.position + 1
	return l
}

func (l *Lexer) ReadChar() byte {
	b := l.b[l.position]
	l.next()
	return b
}

func (l *Lexer) HasNext() bool {
	return l.readPosition < len(l.b)
}

func (l *Lexer) next() {
	l.position++
	l.readPosition = l.position + 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	b := l.ReadChar()

	switch b {
	case '\\':
		b = l.ReadChar()
		switch b {
		case 'd':
			tok.Literal = "d"
			tok.Type = token.DIGIT
		default:
			panic("undefined character:" + string(b))
		}
	case '{':
		tok.Literal = "{"
		tok.Type = token.LBRACE
	case '}':
		tok.Literal = "}"
		tok.Type = token.RBRACE
	case 0, '\n':
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// TODO: 文字列読み込み
		if isDigit(b) {
			// TODO: 複数の数字
			tok.Literal = string(b)
			tok.Type = token.INT

		} else {
			panic("undefined character:" + string(b))
		}
	}

	return tok
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
