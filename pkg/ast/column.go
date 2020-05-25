package ast

import (
	"fmt"

	"github.com/tomoyamachi/dbscheme2struct/pkg/token"
)

type ColumnDef struct {
	Name string
	Type token.Token
	// Elems is the element list for enum and set type.
	Elems   []string
	Options []*ColumnOption
}

func (c *ColumnDef) GetGoType() (*token.GoType, error) {
	gotypes, ok := token.DataTypesGoTypes[c.Type.Literal]
	if !ok {
		return nil, fmt.Errorf("unsupported column type: %s", c.Type.Literal)
	}
	if c.ContainsOpt(ColumnOptionNotNull) {
		return &gotypes.Normal, nil
	}
	return &gotypes.Null, nil
}

func (c *ColumnDef) ContainsOpt(t ColumnOptionType) bool {
	for _, opt := range c.Options {
		if opt.Type == t {
			return true
		}
	}
	return false
}
