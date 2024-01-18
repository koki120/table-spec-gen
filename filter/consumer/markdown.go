package consumer

import (
	"html/template"
	"io"

	"github.com/koki120/table-spec-gen/pipe"
)

func ExportToMarkdown(output io.Writer, tables []pipe.Table) error {
	markdownTemplate := `
# Table Specification
{{range .}}
## {{.TableName}}
| Name | Type | Nullable | Constraints | Referenced | Default | Extra |
|-------------|----------------|-------------|-------------|-------|------------------------|-------------------|
{{range .Columns}}| {{.ColumnName}} | {{.ColumnType}} | {{.IsNullable}} | {{.ConstraintTypes}} | [{{.ReferencedTableName}}](#{{.ReferencedTableName}}) | {{.ColumnDefault}} | {{.Extra}} |
{{end}}
{{end}}
`

	tmpl, err := template.New("tableTemplate").Parse(markdownTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(output, tables)

}
