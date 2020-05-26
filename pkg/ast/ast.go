package ast

import (
	"fmt"

	"github.com/tomoyamachi/ora2struct/pkg/token"
)

type Ddl interface {
	CreateOutput() (Imports, *OutputTable, error)
}

type SQL struct {
	Ddls []Ddl
}

type CreateTable struct {
	Token       token.Token
	Table       TableName
	Columns     []*ColumnDef
	Constraints []*Constraint
}

func (n CreateTable) CreateOutput() (Imports, *OutputTable, error) {
	imports := []string{}
	cols := []ColumnParam{}
	for _, col := range n.Columns {
		gotype, err := col.GetGoType()
		if err != nil {
			return nil, nil, fmt.Errorf("get gotype %s.%s: %w", n.Table.Table, col.Name, err)
		}
		cols = append(cols, ColumnParam{
			Name:   col.Name,
			Type:   gotype.Type,
			Origin: col.Type.Literal,
		})
		imports = append(imports, gotype.Imports...)
	}
	return imports, &OutputTable{Table: n.Table, Type: TypeTable, Columns: cols}, nil
}

type CreateView struct {
	Token       token.Token
	Table       TableName
	Columns     []*ColumnDef
	Constraints []*Constraint
}

func (n CreateView) CreateOutput() (Imports, *OutputTable, error) {
	imports := []string{}
	cols := []ColumnParam{}
	for _, col := range n.Columns {
		cols = append(cols, ColumnParam{
			Name:   col.Name,
			Type:   "interface{}",
			Origin: col.Type.Literal,
		})
	}
	return imports, &OutputTable{Table: n.Table, Type: TypeView, Columns: cols}, nil
}

// ColumnOption is used for parsing column constraint info from SQL.
type ColumnOption struct {
	Type  ColumnOptionType
	Value string
}

// ColumnOptionType is the type for ColumnOption.
type ColumnOptionType int

// ConstraintType is the type for Constraint.
type ConstraintType int

// options
const (
	ColumnOptionNotNull  ColumnOptionType = iota
	ConstraintForeignKey ConstraintType   = iota
)

// Constraint is constraint for table definition.
type Constraint struct {
	Type ConstraintType
	Name string
	// Used for PRIMARY KEY, UNIQUE, ......
	Keys []string
}
