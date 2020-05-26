package parser

import (
	"fmt"
	"log"

	"github.com/tomoyamachi/ora2struct/pkg/ast"
	"github.com/tomoyamachi/ora2struct/pkg/lexer"
	"github.com/tomoyamachi/ora2struct/pkg/token"
)

type (
	prefixParseFn func() ast.Ddl
	infixParseFn  func(node ast.Ddl) ast.Ddl
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.CREATE, p.parseCreateExpression)
	return p
}

func (p *Parser) ParseSQL(debug bool) []ast.Ddl {
	nodes := []ast.Ddl{}

	for p.curToken.Type != token.EOF {
		node := p.parseNode()
		if node != nil {
			nodes = append(nodes, node)
		}
		p.nextToken()
	}
	if debug {
		for _, e := range p.errors {
			log.Println(e)
		}
	}
	return nodes
}

func (p *Parser) parseNode() ast.Ddl {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		// p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.SEMICOLON) {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) curTokensAre(ts ...token.TokenType) bool {
	for _, t := range ts {
		if p.curToken.Type == t {
			return true
		}
	}
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekTokensAre(ts ...token.TokenType) bool {
	for _, t := range ts {
		if p.peekToken.Type == t {
			return true
		}
	}
	return false
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) expectPeeks(ts ...token.TokenType) bool {
	for _, t := range ts {
		if p.peekTokenIs(t) {
			p.nextToken()
			return true
		}
	}
	p.peekError(ts...)
	return false
}

func (p *Parser) peekError(ts ...token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %v, got %s instead", ts, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
