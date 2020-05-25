package generator

const simpleStruct = `
package {{.Package}}

{{if .Imports}}
import (
{{range $import := .Imports}}	"{{$import}}"
{{end}})
{{end}}

{{range .Tables -}}
// Table {{.Table.User}}.{{.Table.Table}}
type {{.Table.Table | ToCamel}} struct {
{{range $column := .Columns}}	{{$column.Name | ToCamel}} {{$column.Type}} ` + "`" + `db:"{{$column.Name}}"` + "`" + ` // type: {{$column.Type}}
{{end}}}

{{end}}
`
