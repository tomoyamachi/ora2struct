package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	STRING = "STRING"

	// Delimiters
	DOT       = "."
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRAKET   = "["
	RBRAKET   = "]"
	COLON     = ":"

	// DataTypes https://docs.oracle.com/cd/E11882_01/server.112/e41085/sqlqr06002.htm#SQLQR959
	VARCHAR2      = "VARCHAR2"
	NVARCHAR2     = "NVARCHAR2"
	NUMBER        = "NUMBER"
	FLOAT         = "FLOAT"
	LONG          = "LONG"
	DATE          = "DATE"
	BINARY_FLOAT  = "BINARY_FLOAT"
	BINARY_DOUBLE = "BINARY_DOUBLE"
	TIMESTAMP     = "TIMESTAMP"
	RAW           = "RAW"
	ROWID         = "ROWID"
	UROWID        = "UROWID"
	CHAR          = "CHAR"
	NCHAR         = "NCHAR"
	CLOB          = "CLOB"
	NCLOB         = "NCLOB"
	BLOB          = "BLOB"
	BFILE         = "BFILE"

	// following datatype contains space
	// TIMESTAMP_WITH_TIME_ZONE
	// INTERVAL_YEAR
	// INTERVAL_DAY
	// LONG_RAW

	CREATE      = "CREATE"
	TABLE       = "TABLE"
	FUNCTION    = "FUNCTION"
	OR          = "OR"
	REPLACE     = "REPLACE"
	EDITIONABLE = "EDITIONABLE"
	VIEW        = "VIEW"
	NOT         = "NOT"
	NULL        = "NULL"
	DEFAULT     = "DEFAULT"
)

var combinations = [][]string{
	{CREATE, TABLE},
	{CREATE, FUNCTION},
	{CREATE, VIEW},
	{CREATE, OR, REPLACE, VIEW},
	{CREATE, OR, REPLACE, EDITIONABLE, VIEW},
	{NOT, NULL},
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	CREATE:        CREATE,
	TABLE:         TABLE,
	VARCHAR2:      VARCHAR2,
	NVARCHAR2:     NVARCHAR2,
	NUMBER:        NUMBER,
	FLOAT:         FLOAT,
	LONG:          LONG,
	DATE:          DATE,
	BINARY_FLOAT:  BINARY_FLOAT,
	BINARY_DOUBLE: BINARY_DOUBLE,
	TIMESTAMP:     TIMESTAMP,
	RAW:           RAW,
	ROWID:         ROWID,
	UROWID:        UROWID,
	CHAR:          CHAR,
	NCHAR:         NCHAR,
	CLOB:          CLOB,
	NCLOB:         NCLOB,
	BLOB:          BLOB,
	BFILE:         BFILE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
