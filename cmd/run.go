package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/tomoyamachi/dbscheme2struct/pkg/ast"
	"github.com/tomoyamachi/dbscheme2struct/pkg/generator"
	"github.com/tomoyamachi/dbscheme2struct/pkg/lexer"
	"github.com/tomoyamachi/dbscheme2struct/pkg/parser"
)

var conf *Config

func NewApp(version string) *cli.App {
	cli.AppHelpTemplate = `NAME:
  {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}
USAGE:
  {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}
VERSION:
  {{.Version}}{{end}}{{end}}{{if .Description}}
DESCRIPTION:
  {{.Description}}{{end}}{{if len .Authors}}
AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
  {{range $index, $author := .Authors}}{{if $index}}
  {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}
OPTIONS:
  {{range $index, $option := .VisibleFlags}}{{if $index}}
  {{end}}{{$option}}{{end}}{{end}}
`
	app := cli.NewApp()
	app.Name = "oracle2struct"
	app.Version = version
	app.Usage = "Output Oracle Database table structures from DDL file"
	app.ArgsUsage = "DDL filename"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		&templateFlag,
		&outputFlag,
		&exportPackageFlag,
		&debugFlag,
	}
	app.Action = run
	return app
}

func run(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return fmt.Errorf("argument length should 1 but got %d", ctx.Args().Len())
	}

	conf = LoadConf(ctx)

	fileName := ctx.Args().First()
	nodes, err := parseFile(fileName)
	if err != nil {
		return fmt.Errorf("parseFile : %w", err)
	}

	output := os.Stdout
	if conf.OutputFile != "" {
		if output, err = os.Create(conf.OutputFile); err != nil {
			return fmt.Errorf("create file %s: %w", conf.OutputFile, err)
		}
	}

	if err = generator.Output(output, nodes, conf.PackageName, conf.TemplateFile); err != nil {
		return fmt.Errorf("output :%w", err)
	}
	return nil
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
	}
	l := lexer.New(stmt)
	p := parser.New(l)
	nodes = p.ParseSQL(conf.Debug)
	return nodes, nil
}
