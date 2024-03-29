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
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.INT, p.parseNumber)
	p.registerPrefix(token.DIGIT, p.parseDigit)
	p.registerPrefix(token.NOTDIGIT, p.parseNotDigit)
	p.registerPrefix(token.STRING, p.parseString)
	p.registerPrefix(token.WORD, p.parseWord)
	p.registerPrefix(token.NOTWORD, p.parseNotWord)
	p.registerPrefix(token.SPACE, p.parseSpace)
	p.registerPrefix(token.NOTSPACE, p.parseNotSpace)
	p.registerPrefix(token.NEWLINE, p.parseNewline)
	p.registerPrefix(token.TAB, p.parseTab)
	p.registerPrefix(token.BACKSLASH, p.parseBackslash)
	p.registerPrefix(token.DOT, p.parseDot)
	p.registerPrefix(token.LPAREN, p.parseBlock)
	p.registerPrefix(token.LBRACKET, p.parseBracket)

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
	return p.parse(token.EOF)
}

func (p *Parser) parse(end token.TokenType) *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(end) {
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

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.curToken,
	}
	stmt.Expression = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseNumber() ast.Expression {
	expression := &ast.NumberExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
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

func (p *Parser) parseNotDigit() ast.Expression {
	expression := &ast.NotDigitExpression{
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

func (p *Parser) parseNotWord() ast.Expression {
	expression := &ast.NotWordExpression{
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

func (p *Parser) parseNotSpace() ast.Expression {
	expression := &ast.NotSpaceExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseNewline() ast.Expression {
	expression := &ast.NewlineExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseTab() ast.Expression {
	expression := &ast.TabExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseBackslash() ast.Expression {
	expression := &ast.BackslashExpression{
		Token: p.curToken,
	}
	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseDot() ast.Expression {
	expression := &ast.DotExpression{
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

func (p *Parser) parseBlock() ast.Expression {
	expression := &ast.BlockExpression{
		Token: p.curToken,
	}
	p.nextToken()

	b := p.parse(token.RPAREN)
	expression.Block = *b

	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) parseBracket() ast.Expression {
	expression := &ast.BracketExpression{
		Token: p.curToken,
	}
	p.nextToken()

	program := &ast.Program{}
	program.Statements = []ast.Statement{}

LoopToken:
	for !p.curTokenIs(token.RBRACKET) {
		switch p.curToken.Type {
		case token.INT, token.STRING:
			var stmt ast.Statement
			for _, stmt := range p.divideToChars(p.curToken) {
				program.Statements = append(program.Statements, stmt)
			}
			if stmt != nil {
				program.Statements = append(program.Statements, stmt)
			}
			p.nextToken()

		case token.LPAREN, token.RPAREN: // do nothing
			p.nextToken()
			continue LoopToken

		case token.LBRACE, token.RBRACE, token.DOT, token.PLUS, token.ASTERISK, token.COMMA:
			stmt := &ast.ExpressionStatement{Token: p.curToken}
			stmt.Expression = &ast.StringExpression{
				Token: p.curToken,
				Range: ast.Range{Min: 1, Max: 1},
			}
			if stmt != nil {
				program.Statements = append(program.Statements, stmt)
			}
			p.nextToken()

		default:
			stmt := p.parseStatement()
			if stmt != nil {
				program.Statements = append(program.Statements, stmt)
			}
			p.nextToken()
		}
	}

	expression.Block = *program

	r, err := p.parseRange()
	if err != nil {
		return nil
	}
	expression.Range = r

	return expression
}

func (p *Parser) divideToChars(t token.Token) []ast.Statement {
	if len(t.Literal) == 0 {
		return nil
	}

	stmts := []ast.Statement{}

	for i := 0; i < len(t.Literal); i++ {
		stmt := &ast.ExpressionStatement{Token: t}

		switch t.Type {
		case token.INT:
			stmt.Expression = &ast.NumberExpression{
				Token: token.Token{
					Type:    token.INT,
					Literal: string(t.Literal[i]),
				},
				Range: ast.Range{Min: 1, Max: 1},
			}
		case token.STRING:
			stmt.Expression = &ast.StringExpression{
				Token: token.Token{
					Type:    token.STRING,
					Literal: string(t.Literal[i]),
				},
				Range: ast.Range{Min: 1, Max: 1},
			}
		default:
			panic("unreachable")
		}

		stmts = append(stmts, stmt)
	}

	return stmts
}
