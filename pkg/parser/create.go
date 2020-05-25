package parser

import (
	"github.com/tomoyamachi/dbscheme2struct/pkg/ast"
	"github.com/tomoyamachi/dbscheme2struct/pkg/token"
)

func (p *Parser) parseCreateExpression() ast.Node {
	switch p.peekToken.Type {
	case token.TABLE:
		return p.parseCreateTable()
	case token.VIEW:
		p.errors = append(p.errors, "CREATE VIEW will support, but currently not")
		return nil
	}
	p.errors = append(p.errors, "CREATE %s is not support", p.peekToken.Literal)
	return nil
}

func (p *Parser) parseCreateTable() ast.Node {
	p.nextToken()
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

func (p *Parser) parseColumns() []*ast.ColumnDef {
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
			for {
				if p.peekTokenIs(token.RPAREN) {
					break
				}
				p.nextToken()
			}
		}

		col.Options = p.parseColumnOpts()
		columns = append(columns, col)
		for {
			p.nextToken()
			if p.curTokenIs(token.COMMA) || p.peekTokenIs(token.RPAREN) || p.peekTokenIs(token.EOF) {
				break
			}

		}
	}
	return columns
}

func (p *Parser) parseColumnOpts() []*ast.ColumnOption {
	opts := []*ast.ColumnOption{}
	for {
		if p.peekTokenIs(token.COMMA) || p.peekTokenIs(token.RPAREN) {
			return opts
		}
		switch p.peekToken.Type {
		case token.NOT:
			p.nextToken()
			if p.expectPeek(token.NULL) {
				opts = append(opts, &ast.ColumnOption{
					Type: ast.ColumnOptionNotNull,
				})
			}
		}
	}
	return nil
}
