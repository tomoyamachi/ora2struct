package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/tomoyamachi/dbscheme2struct/pkg/ast"
	"github.com/tomoyamachi/dbscheme2struct/pkg/lexer"
	"github.com/tomoyamachi/dbscheme2struct/pkg/output"
	"github.com/tomoyamachi/dbscheme2struct/pkg/parser"
)

func main() {
	fileName := os.Args[1]
	nodes, err := parseFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = output.Output(nodes, "db", "")
	if err != nil {
		log.Fatal(err)
	}
}

func parseFile(fileName string) ([]ast.Node, error) {
	ddl, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer ddl.Close()

	var nodes []ast.Node
	scanner := bufio.NewScanner(ddl)
	stmt := ""
	for scanner.Scan() {
		line := scanner.Text()
		stmt += " " + line
		if strings.Contains(line, ";") {
			l := lexer.New(stmt)
			p := parser.New(l)
			nodes = append(p.ParseSQL(), nodes...)
			stmt = ""
		}
	}
	return nodes, nil
}
