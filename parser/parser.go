package parser

import (
	"errors"
	"github.com/cipepser/stringRandom/ast"
	"github.com/cipepser/stringRandom/lexer"
	"github.com/cipepser/stringRandom/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	// TODO: infixになることある？
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.DIGIT, p.parseDigit)
	p.registerPrefix(token.STRING, p.parseString)
	p.registerPrefix(token.WORD, p.parseWord)
	p.registerPrefix(token.SPACE, p.parseSpace)

	p.nextToken()
	p.nextToken()

	return p
}

type prefixParseFn func() ast.Expression

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	//fmt.Println("[DEBUG] program:", program)
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.curToken,
	}
	stmt.Expression = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseDigit() ast.Expression {
	expression := &ast.DigitExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseString() ast.Expression {
	expression := &ast.StringExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseWord() ast.Expression {
	expression := &ast.WordExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseSpace() ast.Expression {
	expression := &ast.SpaceExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseRange() (ast.Range, error) {
	r := ast.Range{}

	switch p.peekToken.Type {
	case token.LBRACE:
		// TODO: LPARENを増やす
		// TODO: カッコごとに対応が必要
		p.nextToken()
		if !p.expectPeek(token.INT) {
			return ast.Range{}, errors.New("unexpected token")
		}

		min, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			return ast.Range{}, errors.New("unexpected token")
		}
		r.Min = min

		switch p.peekToken.Type {
		case token.RBRACE:
			r.Max = min
			p.nextToken()
		case token.COMMA:
			p.nextToken()
			if !p.expectPeek(token.INT) {
				return ast.Range{}, errors.New("unexpected token")
			}

			max, err := strconv.Atoi(p.curToken.Literal)
			if err != nil {
				return ast.Range{}, err
			}
			r.Max = max

			if !p.expectPeek(token.RBRACE) {
				return ast.Range{}, errors.New("unexpected token")
			}
		default:
			panic("unreachable")
		}

	case token.PLUS:
		r.Min = 1
		r.Max = ast.INFINITE
		p.nextToken()

	case token.ASTERISK:
		r.Min = 0
		r.Max = ast.INFINITE
		p.nextToken()

	default:
		r.Min = 1
		r.Max = 1
	}

	return r, nil
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	// TODO: !p.peekTokenIsのループいる？

	return leftExp
}
