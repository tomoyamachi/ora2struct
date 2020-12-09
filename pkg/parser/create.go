package parser

import (
	"fmt"
	"strconv"

	"github.com/tomoyamachi/ora2struct/pkg/ast"
	"github.com/tomoyamachi/ora2struct/pkg/token"
)

func (p *Parser) parseCreateExpression() ast.Ddl {
	for {
		switch p.peekToken.Type {
		case token.TABLE:
			return p.parseCreateTable()
		case token.VIEW:
			return p.parseCreateView()
		case token.FUNCTION:
			p.errors = append(p.errors, "unsupport CREATE FUNCTION")
			return nil
		case token.SEMICOLON, token.EOF:
			p.errors = append(p.errors, "CREATE statement found, but not found resource")
			return nil
		default:
			p.nextToken()
		}
	}
}

func (p *Parser) parseCreateView() ast.Ddl {
	p.nextToken()
	ddl := &ast.CreateView{
		Token: p.curToken,
	}
	ddl.Table = p.parseTableName()
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	ddl.Columns = p.parseViewColumns()
	return ddl
}

func (p *Parser) parseViewColumns() []*ast.ColumnDef {
	columns := []*ast.ColumnDef{}
	for {
		if !p.peekTokensAre(token.STRING, token.IDENT) {
			return columns
		}
		cName := p.peekToken.Literal
		col := &ast.ColumnDef{
			Name: cName,
			Type: token.Token{
				Type:    "interface{}",
				Literal: p.peekToken.Literal,
			},
		}
		columns = append(columns, col)
		for {
			p.nextToken()
			if p.curTokenIs(token.COMMA) {
				break
			}
			if p.curTokenIs(token.LPAREN) {
				p.skipToRParen()
			}
			if p.curTokensAre(token.RPAREN, token.EOF) {
				return columns
			}
		}
	}
}

func (p *Parser) skipToRParen() {
	for {
		p.nextToken()
		if p.curTokenIs(token.RPAREN) {
			return
		}
	}
}

func (p *Parser) parseCreateTable() ast.Ddl {
	p.nextToken()
	ddl := &ast.CreateTable{
		Token: p.curToken,
	}
	ddl.Table = p.parseTableName()
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	ddl.Columns = p.parseTableColumns()
	for {
		p.nextToken()
		if p.curTokenIs(token.EOF) || p.curTokenIs(token.RPAREN) || p.curTokenIs(token.SEMICOLON) {
			break
		}
	}
	return ddl
}

func (p *Parser) parseTableName() ast.TableName {
	if !p.expectPeeks(token.STRING, token.IDENT) {
		return ast.TableName{}
	}
	user := p.curToken.Literal
	if !p.expectPeek(token.DOT) {
		return ast.TableName{}
	}
	if !p.expectPeeks(token.STRING, token.IDENT) {
		return ast.TableName{}
	}
	table := p.curToken.Literal
	return ast.TableName{
		User:  user,
		Table: table,
	}
}

func (p *Parser) parseTableColumns() []*ast.ColumnDef {
	columns := []*ast.ColumnDef{}
	for {
		if !p.expectPeeks(token.STRING, token.IDENT) {
			return columns
		}
		cName := p.curToken.Literal
		if _, ok := token.DataTypesGoTypes[p.peekToken.Literal]; !ok {
			p.peekError(p.peekToken.Type)
			return columns
		}
		p.nextToken()
		// add new columns

		col := &ast.ColumnDef{
			Name:    cName,
			Type:    p.curToken,
			Elems:   nil,
			Options: nil,
		}

		// some columns need bracket options, like VARCHAR2(30)
		if p.peekTokenIs(token.LPAREN) {
			// if number checks paren value
			if col.Type.Type == token.NUMBER {
				if i, err := p.parseNumberScale(); err != nil {
					p.errors = append(p.errors, fmt.Sprintf("parse NUMBER scale column: %s, err: %s", cName, err))
				} else if i > 0 {
					col.Type = token.Token{token.FLOAT, token.FLOAT}
				}
			}
			p.skipToRParen()
		}

		col.Options = p.parseColumnOpts()
		columns = append(columns, col)
		for {
			// skip to last of a column definition
			p.nextToken()
			if p.curTokenIs(token.COMMA) || p.peekTokenIs(token.RPAREN) || p.peekTokenIs(token.EOF) {
				break
			}

		}
	}
	return columns
}

func (p *Parser) parseNumberScale() (int64, error) {
	p.nextToken() // move to LPAREN
	p.nextToken() // move to precision
	if p.peekTokenIs(token.RPAREN) {
		return 0, nil
	}

	if !p.peekTokenIs(token.COMMA) {
		return 0, fmt.Errorf("expect comma but got %s", p.peekToken.Literal)
	}
	p.nextToken()
	return strconv.ParseInt(p.peekToken.Literal, 10, 64)
}

func (p *Parser) parseColumnOpts() []*ast.ColumnOption {
	opts := []*ast.ColumnOption{}
	for {
		if p.peekTokenIs(token.LPAREN) {
			p.skipToRParen()
		}
		if p.peekTokenIs(token.COMMA) || p.peekTokenIs(token.RPAREN) {
			return opts
		}
		switch p.curToken.Type {
		case token.NOT:
			if p.peekTokenIs(token.NULL) {
				opts = append(opts, &ast.ColumnOption{
					Type: ast.ColumnOptionNotNull,
				})
			}
		}
		p.nextToken()
	}
	return opts
}
