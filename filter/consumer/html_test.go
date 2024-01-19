package consumer_test

import (
	"io"
	"os"
	"testing"

	"github.com/koki120/table-spec-gen/filter/consumer"
	"github.com/koki120/table-spec-gen/pipe"
)

func TestExportToHTML(t *testing.T) {

	type args struct {
		output io.Writer
		tables []pipe.TableMetaData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				output: os.Stdout,
				tables: []pipe.TableMetaData{
					{
						TableName: "table1",
						Columns: []pipe.Column{
							{
								ColumnName:          "ColumnName1",
								ColumnDefault:       "ColumnDefault1",
								IsNullable:          "IsNullable1",
								ColumnType:          "ColumnType1",
								Extra:               "Extra1",
								ReferencedTableName: "ReferencedTableName1",
								ConstraintTypes:     "ConstraintTypes1",
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			if err := consumer.ExportToHTML(tt.args.output, tt.args.tables); (err != nil) == tt.wantErr {
				t.Errorf("ExportToHTML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
