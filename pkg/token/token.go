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

	CREATE = "CREATE"
	TABLE  = "TABLE"
)

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

var DataTypesGoType = map[string]string{
	VARCHAR2:      "string",
	NVARCHAR2:     "string",
	NUMBER:        "int",
	FLOAT:         "float32",
	LONG:          "string",
	DATE:          "time.Time",
	BINARY_FLOAT:  "[]byte",
	BINARY_DOUBLE: "[]byte",
	TIMESTAMP:     "int",
	RAW:           "[]byte",
	ROWID:         "string",
	UROWID:        "string",
	CHAR:          "string",
	NCHAR:         "string",
	CLOB:          "[]byte",
	NCLOB:         "[]byte",
	BLOB:          "[]byte",
	BFILE:         "[]byte",
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
