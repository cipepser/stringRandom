package ast

import (
	"bytes"
	"github.com/cipepser/stringRandom/token"
	"math"
	"strconv"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type DigitExpression struct {
	Token token.Token
	Range
}

func (de *DigitExpression) TokenLiteral() string { return de.Token.Literal }
func (de *DigitExpression) String() string {
	var out bytes.Buffer

	out.WriteString("\\d")
	switch de.Range.max {
	case INFINITE:
		if de.Range.min == 0 {
			out.WriteString("*")
		} else {
			out.WriteString("+")
		}
	default:
		out.WriteString("{")
		out.WriteString(strconv.Itoa(de.Range.min))
		out.WriteString(",")
		out.WriteString(strconv.Itoa(de.Range.max))
		out.WriteString("}")
	}

	return out.String()
}
func (de *DigitExpression) expressionNode() {}

type Range struct {
	min, max int
}

const (
	INFINITE = math.MaxInt64 // 便宜上MaxInt64を使う
)
