package lexer

import (
	"github.com/cipepser/stringRandom/token"
)

type Lexer struct {
	b            string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{b: input}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.b) {
		l.ch = 0
	} else {
		l.ch = l.b[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) HasNext() bool {
	return l.readPosition < len(l.b)
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '\\':
		switch l.peekChar() {
		case 'd':
			l.ReadChar()
			tok.Literal = "d"
			tok.Type = token.DIGIT
		default:
			panic("undefined character(escape):" + string(l.ch))
		}
	case '+':
		tok.Literal = "+"
		tok.Type = token.PLUS
	case '*':
		tok.Literal = "*"
		tok.Type = token.ASTERISK
	case '{':
		tok.Literal = "{"
		tok.Type = token.LBRACE
	case '}':
		tok.Literal = "}"
		tok.Type = token.RBRACE
	case ',':
		tok.Literal = ","
		tok.Type = token.COMMA
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// TODO: 文字列読み込み
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			panic("undefined character(default):" + string(l.ch))
		}
	}

	l.ReadChar()
	return tok
}

// TODO: isLetterの実装

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.ReadChar()
	}
	return l.b[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.b) {
		return 0
	} else {
		return l.b[l.readPosition]
	}
}
