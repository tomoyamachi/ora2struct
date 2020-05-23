package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tomoyamachi/dbscheme2struct/pkg/ast"
	"github.com/tomoyamachi/dbscheme2struct/pkg/lexer"
	"github.com/tomoyamachi/dbscheme2struct/pkg/parser"
)

func main() {
	ddl, err := os.Open("./test/ddl/single.ddl")
	if err != nil {
		log.Fatal(err)
	}

	var nodes []ast.Node
	scanner := bufio.NewScanner(ddl)
	stmt := ""
	for scanner.Scan() {
		line := scanner.Text()
		stmt += " " + line
		if strings.Contains(line, ";") {
			fmt.Println(stmt)
			l := lexer.New(stmt)
			p := parser.New(l)
			nodes = p.ParseSQL()
			stmt = ""
		}
	}
	fmt.Println(nodes)
}
