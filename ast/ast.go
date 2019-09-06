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

type StringExpression struct {
	Token token.Token
	Range
}

func (se *StringExpression) TokenLiteral() string { return se.Token.Literal }
func (se *StringExpression) String() string {
	var out bytes.Buffer

	out.WriteString(se.Token.Literal)
	if se.Range.Min == 0 && se.Range.Max == 0 {
		return out.String()
	}

	switch se.Range.Max {
	case INFINITE:
		if se.Range.Min == 0 {
			out.WriteString("*")
		} else {
			out.WriteString("+")
		}
	default:
		out.WriteString("{")
		out.WriteString(strconv.Itoa(se.Range.Min))
		out.WriteString(",")
		out.WriteString(strconv.Itoa(se.Range.Max))
		out.WriteString("}")
	}

	return out.String()
}
func (se *StringExpression) expressionNode() {}

type WordExpression struct {
	Token token.Token
	Range
}

func (we *WordExpression) TokenLiteral() string { return we.Token.Literal }
func (we *WordExpression) String() string {
	var out bytes.Buffer

	out.WriteString("\\w")
	switch we.Range.Max {
	case INFINITE:
		if we.Range.Min == 0 {
			out.WriteString("*")
		} else {
			out.WriteString("+")
		}
	default:
		out.WriteString("{")
		out.WriteString(strconv.Itoa(we.Range.Min))
		out.WriteString(",")
		out.WriteString(strconv.Itoa(we.Range.Max))
		out.WriteString("}")
	}

	return out.String()
}
func (we *WordExpression) expressionNode() {}

type SpaceExpression struct {
	Token token.Token
	Range
}

func (se *SpaceExpression) TokenLiteral() string { return se.Token.Literal }
func (se *SpaceExpression) String() string {
	var out bytes.Buffer

	out.WriteString("\\s")
	switch se.Range.Max {
	case INFINITE:
		if se.Range.Min == 0 {
			out.WriteString("*")
		} else {
			out.WriteString("+")
		}
	default:
		out.WriteString("{")
		out.WriteString(strconv.Itoa(se.Range.Min))
		out.WriteString(",")
		out.WriteString(strconv.Itoa(se.Range.Max))
		out.WriteString("}")
	}

	return out.String()
}
func (se *SpaceExpression) expressionNode() {}
