package generator

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/tomoyamachi/ora2struct/pkg/ast"
)

type tplData struct {
	Package string
	Tables  []*ast.OutputTable
	Imports []string
}

func Output(writer io.Writer, nodes []ast.Ddl, pkgName, tplFile string) (err error) {
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

func convertOutputTpl(nodes []ast.Ddl) (tplData, error) {
	imports := []string{}
	tables := []*ast.OutputTable{}
	for _, node := range nodes {
		impt, tbl, err := node.CreateOutput()
		if err != nil {
			return tplData{}, fmt.Errorf("convertTable :%w", err)
		}
		imports = append(imports, impt...)
		tables = append(tables, tbl)
	}
	return tplData{
		Tables:  tables,
		Imports: unique(imports),
	}, nil
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
