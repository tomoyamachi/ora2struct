package generator

const simpleStruct = `
package {{.Package}}
{{range .Nodes -}}
// Table {{.Table.User}}.{{.Table.Table}}
type {{.Table.Table | ToCamel}} struct {
{{range $column := .Columns}} {{$column.Name | ToCamel}} {{$column.Type.Literal | ToGoType}} ` + "`" + `db:"{{$column.Name}}"` + "`" + ` // type: {{$column.Type.Literal}}
{{end}}}

{{end}}
`
