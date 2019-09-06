package lexer

import (
	"github.com/cipepser/stringRandom/token"
	"strings"
)

var (
	RANGECHARS = "*+{"
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
			panic("undefined meta-character:" + string(l.ch))
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
		// TODO: 次の文字が範囲指定文字（*とか+とか{）だった場合に失敗する
		if isLetter(l.ch) {
			tok.Literal = l.readString()
			tok.Type = token.STRING
			return tok
		} else if isDigit(l.ch) {
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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
	// TODO: これ以外の文字も追加する
}

func (l *Lexer) readString() string {
	position := l.position
	for isLetter(l.ch) {
		if strings.Contains(RANGECHARS, string(l.peekChar())) && position != l.position {
			return l.b[position:l.position]
		}
		l.ReadChar()
	}
	return l.b[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		if strings.Contains(RANGECHARS, string(l.peekChar())) && position != l.position {
			return l.b[position:l.position]
		}
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
