package generate

import (
	"text/template"
)

// FileData contains top-level infos for templates
type FileData struct {
	Package string
}

// FieldData describes a field to be mapped
type FieldData struct {
	Name       string
	Type       string
	ColumnName string
	ColumnType string
}

// StructData describes a struct to be mapped
type StructData struct {
	Name      string
	TableName string
	Fields    []FieldData
}

var (
	prologTemplate = template.Must(template.New("prolog").Parse(
		`// generated with yagorm
package {{ .Package }}

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/aacanakin/qb"

	"bitbucket.org/cdevienne/yagorm"
)

`))

	structTemplate = template.Must(template.New("struct").Parse(`
var {{ .TableName }}Table = qb.Table(
	"{{ .TableName }}",
	{{- range $_, $fd := .Fields }}
	qb.Column("{{ $fd.ColumnName }}", {{ $fd.ColumnType }}),
	{{- end }}
)
`))
)
