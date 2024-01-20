package transformer_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/koki120/table-spec-gen/config"
	"github.com/koki120/table-spec-gen/filter/producer"
	"github.com/koki120/table-spec-gen/filter/transformer"
	"github.com/koki120/table-spec-gen/pipe"
)

func TestConvertColumnMetadataToTableMetaData(t *testing.T) {

	type args struct {
		columnMetadata []pipe.ColumnMetadata
	}
	tests := []struct {
		name string
		args args
		want []pipe.TableMetaData
	}{
		{
			name: "success",
			args: args{
				columnMetadata: []pipe.ColumnMetadata{
					{
						TableName:           "table1",
						ColumnName:          "ColumnName1",
						ColumnDefault:       "ColumnDefault1",
						IsNullable:          "IsNullable1",
						ColumnType:          "ColumnType1",
						Extra:               "Extra1",
						ReferencedTableName: "ReferencedTableName1",
						ConstraintTypes:     "ConstraintTypes1",
					},
					{
						TableName:           "table1",
						ColumnName:          "ColumnName2",
						ColumnDefault:       "ColumnDefault2",
						IsNullable:          "IsNullable2",
						ColumnType:          "ColumnType2",
						Extra:               "Extra2",
						ReferencedTableName: "ReferencedTableName2",
						ConstraintTypes:     "ConstraintTypes2",
					},
					{
						TableName:           "table2",
						ColumnName:          "ColumnName3",
						ColumnDefault:       "ColumnDefault3",
						IsNullable:          "IsNullable3",
						ColumnType:          "ColumnType3",
						Extra:               "Extra3",
						ReferencedTableName: "ReferencedTableName3",
						ConstraintTypes:     "ConstraintTypes3",
					},
				},
			},
			want: []pipe.TableMetaData{
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
						{
							ColumnName:          "ColumnName2",
							ColumnDefault:       "ColumnDefault2",
							IsNullable:          "IsNullable2",
							ColumnType:          "ColumnType2",
							Extra:               "Extra2",
							ReferencedTableName: "ReferencedTableName2",
							ConstraintTypes:     "ConstraintTypes2",
						},
					},
				},
				{
					TableName: "table2",
					Columns: []pipe.Column{
						{
							ColumnName:          "ColumnName3",
							ColumnDefault:       "ColumnDefault3",
							IsNullable:          "IsNullable3",
							ColumnType:          "ColumnType3",
							Extra:               "Extra3",
							ReferencedTableName: "ReferencedTableName3",
							ConstraintTypes:     "ConstraintTypes3",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := transformer.ConvertColumnMetadataToTableMetaData(tt.args.columnMetadata)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("diff =%v", cmp.Diff(got, tt.want))
			}

		})
	}
}

func TestConvertSQLRowsToTableMetadata(t *testing.T) {

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
			source, err := producer.FetchColumnMetadata(informationSchemaDB, config.DBName())
			if err != nil {
				t.Error(err)
			}
			got, err := transformer.ConvertSQLRowsToTableMetadata(source)
			if (err != nil) != tt.hasErr {
				t.Errorf("ConvertSQLRowsToTableMetadata() error = %v, hasErr %v", err, tt.hasErr)
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("diff =%v", cmp.Diff(got, tt.want))
			}
		})
	}
}
