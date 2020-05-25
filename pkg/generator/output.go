package generator

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/tomoyamachi/dbscheme2struct/pkg/ast"
)

type tplData struct {
	Package string
	Tables  []*OutputTable
	Views   []*OutputView
	Imports []string
}

type OutputTable struct {
	Table   ast.TableName
	Columns []ColumnParam
}

type OutputView struct {
	View    ast.TableName
	Columns []ColumnParam
}

type ColumnParam struct {
	Name   string
	Type   string
	Import string
}

func Output(writer io.Writer, nodes []ast.Node, pkgName, tplFile string) (err error) {
	var tplSource string
	if tplFile == "" {
		tplSource = simpleStruct
	} else {
		tplSource, err = loadFileStr(tplFile)
		if err != nil {
			return err
		}
	}
	funcMap := template.FuncMap{"ToCamel": toCamel}
	tpl, err := template.New("struct").Funcs(funcMap).Parse(tplSource)
	if err != nil {
		return err
	}
	d, err := convertOutputTpl(nodes)
	d.Package = pkgName
	return tpl.Execute(writer, d)
}

func convertOutputTpl(nodes []ast.Node) (tplData, error) {
	imports := []string{}
	tables := []*OutputTable{}
	views := []*OutputView{}
	for _, node := range nodes {
		switch node.(type) {
		case *ast.CreateTable:
			impt, tbl, err := convertTable(node.(*ast.CreateTable))
			if err != nil {
				return tplData{}, fmt.Errorf("convertTable :%w", err)
			}
			imports = append(imports, impt...)
			tables = append(tables, tbl)
		case *ast.CreateView:
			log.Println("view")
		default:
			return tplData{}, fmt.Errorf("unsupport node: %v", node)
		}
	}
	return tplData{
		Tables:  tables,
		Views:   views,
		Imports: unique(imports),
	}, nil
}

func convertTable(t *ast.CreateTable) (imports []string, ot *OutputTable, err error) {
	imports = []string{}
	cols := []ColumnParam{}
	for _, col := range t.Columns {
		gotype, err := col.GetGoType()
		if err != nil {
			return nil, nil, fmt.Errorf("get gotype %s.%s: %w", t.Table.Table, col.Name, err)
		}
		cols = append(cols, ColumnParam{
			Name: col.Name,
			Type: gotype.Type,
		})
		imports = append(imports, gotype.Imports...)
	}
	return imports, &OutputTable{Table: t.Table, Columns: cols}, nil
}

func loadFileStr(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func toCamel(str string) string {
	s := ""
	capitalize := true
	for _, ch := range str {
		if ch == '_' {
			capitalize = true
			continue
		}
		if capitalize {
			s += strings.ToUpper(string(ch))
			capitalize = false
			continue
		}
		s += strings.ToLower(string(ch))
	}
	return s
}

func unique(args []string) []string {
	results := make([]string, 0, len(args))
	dup := map[string]struct{}{}
	for _, s := range args {
		if _, ok := dup[s]; !ok {
			dup[s] = struct{}{}
			results = append(results, s)
		}
	}
	return results
}
