package parser

import (
	"fmt"

	"github.com/tomoyamachi/dbscheme2struct/pkg/ast"
	"github.com/tomoyamachi/dbscheme2struct/pkg/lexer"
	"github.com/tomoyamachi/dbscheme2struct/pkg/token"
)

type (
	prefixParseFn func() ast.Node
	infixParseFn  func(node ast.Node) ast.Node
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

func (p *Parser) ParseSQL() []ast.Node {
	nodes := []ast.Node{}

	for p.curToken.Type != token.EOF {
		node := p.parseNode()
		if node != nil {
			nodes = append(nodes, node)
		}
		p.nextToken()
	}
	return nodes
}

func (p *Parser) parseNode() ast.Node {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
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

func (p *Parser) parseCreateExpression() ast.Node {
	if !p.expectPeek(token.TABLE) {
		fmt.Println("only allow CREATE TABLE, got", p.peekToken)
		return nil
	}
	node := &ast.CreateTable{
		Token: p.curToken,
	}
	node.Table = p.parseTableName()
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	node.Columns = p.parseColumns()
	for {
		p.nextToken()
		if p.curTokenIs(token.EOF) {
			break
		}
	}
	return node
}

func (p *Parser) parseTableName() ast.TableName {
	if !p.expectPeek(token.STRING) {
		return ast.TableName{}
	}
	user := p.curToken.Literal
	if !p.expectPeek(token.DOT) {
		return ast.TableName{}
	}
	if !p.expectPeek(token.STRING) {
		return ast.TableName{}
	}
	table := p.curToken.Literal
	return ast.TableName{
		User:  user,
		Table: table,
	}
}

func (p *Parser) parseColumns() []*ast.ColumnDef {
	columns := []*ast.ColumnDef{}

	for {
		if !p.expectPeek(token.STRING) {
			return columns
		}
		cName := p.curToken.Literal
		if _, ok := token.DataTypes[p.peekToken.Literal]; !ok {
			p.peekError(p.peekToken.Type)
			return columns
		}
		p.nextToken()
		// add new columns
		columns = append(columns, &ast.ColumnDef{
			Name:    cName,
			Type:    p.curToken,
			Elems:   nil,
			Options: nil,
		})

		if p.peekTokenIs(token.LPAREN) {
			// skip to rparen
			// log.Println("skipt to RPAREN")
			for {
				if p.peekTokenIs(token.RPAREN) {
					break
				}
				p.nextToken()
			}
		}

		for {
			p.nextToken()
			if p.curTokenIs(token.COMMA) || p.peekTokenIs(token.RPAREN) || p.peekTokenIs(token.EOF) {
				break
			}

		}
	}
	return columns
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

	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
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
