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

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type DigitExpression struct {
	Token token.Token
	Range
}

func (de *DigitExpression) TokenLiteral() string { return de.Token.Literal }
func (de *DigitExpression) String() string {
	var out bytes.Buffer

	out.WriteString("\\d")
	switch de.Range.Max {
	case INFINITE:
		if de.Range.Min == 0 {
			out.WriteString("*")
		} else {
			out.WriteString("+")
		}
	default:
		out.WriteString("{")
		out.WriteString(strconv.Itoa(de.Range.Min))
		out.WriteString(",")
		out.WriteString(strconv.Itoa(de.Range.Max))
		out.WriteString("}")
	}

	return out.String()
}
func (de *DigitExpression) expressionNode() {}

type Range struct {
	Min, Max int
}

const (
	INFINITE = math.MaxInt64 // 便宜上MaxInt64を使う
)
