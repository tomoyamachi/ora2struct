package output

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/tomoyamachi/dbscheme2struct/pkg/ast"
	"github.com/tomoyamachi/dbscheme2struct/pkg/token"
)

type tplData struct {
	Package string
	Nodes   []ast.Node
}

func Output(nodes []ast.Node, pkgName, tplFile string) (err error) {
	var tplSource string
	if tplFile == "" {
		tplSource = simpleStruct
	} else {
		tplSource, err = loadFileStr(tplFile)
		if err != nil {
			return err
		}
	}
	funcMap := template.FuncMap{"ToCamel": toCamel, "ToGoType": toGoType}
	tpl, err := template.New("struct").Funcs(funcMap).Parse(tplSource)
	if err != nil {
		return err
	}
	return tpl.Execute(os.Stdout, tplData{Nodes: nodes, Package: pkgName})
}

func loadFileStr(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func toGoType(dbtype string) string {
	gotype, ok := token.DataTypesGoType[dbtype]
	if !ok {
		log.Fatal("invalidtype", dbtype)
	}
	return gotype
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
