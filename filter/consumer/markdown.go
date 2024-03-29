package consumer

import (
	"io"
	"text/template"

	"github.com/koki120/table-spec-gen/pipe"
)

func ExportToMarkdown(output io.Writer, tables []pipe.TableMetaData) error {
	markdownTemplate := `
# Table Specification
{{range .}}
## {{.TableName}}
| Name | Type | Nullable | Constraints | Referenced | Default | Extra |
|-------------|----------------|-------------|-------------|-------|------------------------|-------------------|
{{range .Columns}}| {{.ColumnName}} | {{.ColumnType}} | {{.IsNullable}} | {{.ConstraintTypes}} | {{if ne .ReferencedTableName ""}}[{{.ReferencedTableName}}](#{{.ReferencedTableName}}){{end}} | {{.ColumnDefault}} | {{.Extra}} |
{{end}}
{{end}}
`

	tmpl, err := template.New("tableTemplate").Parse(markdownTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(output, tables)

}
