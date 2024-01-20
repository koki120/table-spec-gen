package producer_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/koki120/table-spec-gen/config"
	"github.com/koki120/table-spec-gen/filter/producer"
	"github.com/koki120/table-spec-gen/pipe"
)

func TestFetchColumnMetadata(t *testing.T) {

	tests := []struct {
		name   string
		want   []pipe.ColumnMetadata
		hasErr bool
	}{
		{
			name: "success",
			want: []pipe.ColumnMetadata{
				{
					TableName:           "orders",
					ColumnName:          "id",
					ColumnDefault:       "",
					IsNullable:          "NO",
					ColumnType:          "int",
					Extra:               "auto_increment",
					ReferencedTableName: "",
					ConstraintTypes:     "PRIMARY KEY",
				},
				{
					TableName:           "orders",
					ColumnName:          "product_name",
					ColumnDefault:       "",
					IsNullable:          "NO",
					ColumnType:          "varchar(255)",
					Extra:               "",
					ReferencedTableName: "",
					ConstraintTypes:     "",
				},
				{
					TableName:           "orders",
					ColumnName:          "quantity",
					ColumnDefault:       "1",
					IsNullable:          "YES",
					ColumnType:          "int",
					Extra:               "",
					ReferencedTableName: "",
					ConstraintTypes:     "",
				},
				{
					TableName:           "orders",
					ColumnName:          "user_id",
					ColumnDefault:       "",
					IsNullable:          "YES",
					ColumnType:          "int",
					Extra:               "",
					ReferencedTableName: "users",
					ConstraintTypes:     "FOREIGN KEY",
				},
				{
					TableName:           "users",
					ColumnName:          "email",
					ColumnDefault:       "",
					IsNullable:          "NO",
					ColumnType:          "varchar(255)",
					Extra:               "",
					ReferencedTableName: "",
					ConstraintTypes:     "UNIQUE",
				},
				{
					TableName:           "users",
					ColumnName:          "id",
					ColumnDefault:       "",
					IsNullable:          "NO",
					ColumnType:          "int",
					Extra:               "auto_increment",
					ReferencedTableName: "",
					ConstraintTypes:     "PRIMARY KEY",
				},
				{
					TableName:           "users",
					ColumnName:          "name",
					ColumnDefault:       "",
					IsNullable:          "NO",
					ColumnType:          "varchar(255)",
					Extra:               "",
					ReferencedTableName: "",
					ConstraintTypes:     "",
				},
			},
			hasErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := producer.FetchColumnMetadata(informationSchemaDB, config.DBName())
			if (err != nil) != tt.hasErr {
				t.Errorf("ConvertSQLRowsToTableMetadata() error = %v, hasErr %v", err, tt.hasErr)
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("diff =%v", cmp.Diff(got, tt.want))
			}
		})
	}
}
