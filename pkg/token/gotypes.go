package token

type GoType struct {
	Type    string
	Imports []string
}

type GoTypes struct {
	Normal GoType
	Null   GoType
}

var (
	stringTypes = GoTypes{
		Normal: GoType{
			Type: "string",
		},
		Null: GoType{
			Type:    "sql.NullString",
			Imports: []string{"database/sql"},
		},
	}
	intTypes = GoTypes{
		Normal: GoType{
			Type: "int64",
		},
		Null: GoType{
			Type:    "sql.NullInt64",
			Imports: []string{"database/sql"},
		},
	}
	floatTypes = GoTypes{
		Normal: GoType{
			Type: "float64",
		},
		Null: GoType{
			Type:    "sql.NullFloat64",
			Imports: []string{"database/sql"},
		},
	}
	timeTypes = GoTypes{
		Normal: GoType{
			Type:    "time.Time",
			Imports: []string{"time"},
		},
		Null: GoType{
			Type:    "sql.NullTime",
			Imports: []string{"database/sql"},
		},
	}
	byteTypes = GoTypes{
		Normal: GoType{
			Type: "[]byte",
		},
		Null: GoType{
			Type: "[]byte",
		},
	}
)

var DataTypesGoTypes = map[string]GoTypes{
	VARCHAR2:      stringTypes,
	NVARCHAR2:     stringTypes,
	NUMBER:        intTypes,
	FLOAT:         floatTypes,
	LONG:          stringTypes,
	DATE:          timeTypes,
	BINARY_FLOAT:  byteTypes,
	BINARY_DOUBLE: byteTypes,
	TIMESTAMP:     intTypes,
	RAW:           byteTypes,
	ROWID:         stringTypes,
	UROWID:        stringTypes,
	CHAR:          stringTypes,
	NCHAR:         stringTypes,
	CLOB:          byteTypes,
	NCLOB:         byteTypes,
	BLOB:          byteTypes,
	BFILE:         byteTypes,
}
