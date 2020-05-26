package ast

type TableType int

const (
	TypeTable TableType = iota
	TypeView
)

type TableName struct {
	User  string
	Table string
}

type OutputTable struct {
	Table   TableName
	Type    TableType
	Columns []ColumnParam
}

type ColumnParam struct {
	Name   string
	Type   string
	Origin string
}

type Imports []string
