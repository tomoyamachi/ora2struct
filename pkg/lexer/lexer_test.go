package lexer

import (
	"testing"

	"github.com/tomoyamachi/ora2struct/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `
  CREATE TABLE "USER_A"."TABLE_A" 
   (	"F1" CHAR(5), 
	"F2" DATE DEFAULT sysdate, 
	"F3" VARCHAR2(100) NOT NULL, 
	"F4" NUMBER(16,0), 
	"F5" CHAR(1) DEFAULT '0' 
   );

  create table USER_B.TABLE_B 
   (	F1 CHAR(5),
	F2 DATE DEFAULT sysdate, 
	F3 VARCHAR2(100) NOT NULL, 
	F4 NUMBER(16,0), 
	F5 CHAR(1) DEFAULT '0'
   );
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CREATE, "CREATE"},
		{token.TABLE, "TABLE"},
		{token.STRING, "USER_A"},
		{token.DOT, "."},
		{token.STRING, "TABLE_A"},
		{token.LPAREN, "("},

		{token.STRING, "F1"},
		{token.CHAR, "CHAR"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.COMMA, ","},

		{token.STRING, "F2"},
		{token.DATE, "DATE"},
		{token.DEFAULT, "DEFAULT"},
		{token.IDENT, "sysdate"},
		{token.COMMA, ","},

		{token.STRING, "F3"},
		{token.VARCHAR2, "VARCHAR2"},
		{token.LPAREN, "("},
		{token.INT, "100"},
		{token.RPAREN, ")"},
		{token.NOT, "NOT"},
		{token.NULL, "NULL"},
		{token.COMMA, ","},

		{token.STRING, "F4"},
		{token.NUMBER, "NUMBER"},
		{token.LPAREN, "("},
		{token.INT, "16"},
		{token.COMMA, ","},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.COMMA, ","},

		{token.STRING, "F5"},
		{token.CHAR, "CHAR"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.DEFAULT, "DEFAULT"},
		{token.STRING, "0"},

		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.CREATE, "create"},
		{token.TABLE, "table"},
		{token.IDENT, "USER_B"},
		{token.DOT, "."},
		{token.IDENT, "TABLE_B"},
		{token.LPAREN, "("},

		{token.IDENT, "F1"},
		{token.CHAR, "CHAR"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.COMMA, ","},

		{token.IDENT, "F2"},
		{token.DATE, "DATE"},
		{token.DEFAULT, "DEFAULT"},
		{token.IDENT, "sysdate"},
		{token.COMMA, ","},

		{token.IDENT, "F3"},
		{token.VARCHAR2, "VARCHAR2"},
		{token.LPAREN, "("},
		{token.INT, "100"},
		{token.RPAREN, ")"},
		{token.NOT, "NOT"},
		{token.NULL, "NULL"},
		{token.COMMA, ","},

		{token.IDENT, "F4"},
		{token.NUMBER, "NUMBER"},
		{token.LPAREN, "("},
		{token.INT, "16"},
		{token.COMMA, ","},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.COMMA, ","},

		{token.IDENT, "F5"},
		{token.CHAR, "CHAR"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.DEFAULT, "DEFAULT"},
		{token.STRING, "0"},

		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q(%s)", i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q(%s)", i, tt.expectedLiteral, tok.Literal, tok.Type)
		}
	}
}
