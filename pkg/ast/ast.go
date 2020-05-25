package ast

import (
	"github.com/tomoyamachi/dbscheme2struct/pkg/token"
)

type Node interface{}

type CreateTable struct {
	Token       token.Token
	Table       TableName
	Columns     []*ColumnDef
	Constraints []*Constraint
}

type CreateView struct {
	Token       token.Token
	Table       TableName
	Columns     []*ColumnDef
	Constraints []*Constraint
}

type TableName struct {
	User  string
	Table string
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

// ConstraintType is the type for Constraint.
type ConstraintType int

// TODO: support Foreign key
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
