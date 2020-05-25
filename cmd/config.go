package main

import "github.com/urfave/cli/v2"

var (
	templateFlag = cli.StringFlag{
		Name:    "template",
		Aliases: []string{"t"},
		Usage:   "use template",
	}
	outputFlag = cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Value:   "models.go",
		Usage:   "output file name",
	}
	exportPackageFlag = cli.StringFlag{
		Name:    "package",
		Aliases: []string{"p"},
		Value:   "models",
		Usage:   "export package name",
	}
	debugFlag = cli.BoolFlag{
		Name:    "debug",
		Aliases: []string{"d"},
		Usage:   "debug mode",
	}
)

type Config struct {
	TemplateFile string
	OutputFile   string
	PackageName  string
	Debug        bool
}

func LoadConf(c *cli.Context) *Config {
	return &Config{
		TemplateFile: c.String("template"),
		OutputFile:   c.String("output"),
		PackageName:  c.String("package"),
		Debug:        c.Bool("debug"),
	}
}
