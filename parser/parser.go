package parser

import (
	"github.com/cipepser/stringRandom/ast"
	"github.com/cipepser/stringRandom/lexer"
	"github.com/cipepser/stringRandom/token"
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
	// TODO
	panic("unimplemented")
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
