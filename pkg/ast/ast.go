package ast

import "github.com/tomoyamachi/dbscheme2struct/pkg/token"

type Node interface{}

type CreateTable struct {
	Token       token.Token
	Action      string
	Table       TableName
	Columns     []*ColumnDef
	Constraints []*Constraint
}

// DDL represents a CREATE, ALTER, DROP or RENAME statement.
// Table is set for AlterStr, DropStr, RenameStr.
// NewName is set for AlterStr, CreateStr, RenameStr.
// type DDL struct {
// 	Action string
// 	Table  TableName
// }

type TableName struct {
	User  string
	Table string
}

type ColumnDef struct {
	Name string
	Type token.Token
	// Elems is the element list for enum and set type.
	Elems   []string
	Options []*ColumnOption
}

// ColumnOption is used for parsing column constraint info from SQL.
type ColumnOption struct {
	Type  ColumnOptionType
	Value string
}

// ColumnOptionType is the type for ColumnOption.
type ColumnOptionType int

const (
	ColumnOptionNoOption ColumnOptionType = iota
	ColumnOptionNotNull
	ColumnOptionNull
)

type ConstraintType int

const (
	ConstraintNoConstraint ConstraintType = iota
	ConstraintForeignKey
)

// Constraint is constraint for table definition.
type Constraint struct {
	Type ConstraintType
	Name string
	// Used for PRIMARY KEY, UNIQUE, ......
	Keys []string
}
