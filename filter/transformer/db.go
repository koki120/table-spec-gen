package transformer

import (
	"database/sql"

	"github.com/koki120/table-spec-gen/pipe"
)

func ConvertSQLRowsToTableMetadata(rows *sql.Rows) ([]pipe.ColumnMetadata, error) {
	defer rows.Close()
	result := make([]pipe.ColumnMetadata, 0, 1000)
	var (
		tableName           sql.NullString
		columnName          sql.NullString
		columnDefault       sql.NullString
		isNullable          sql.NullString
		columnType          sql.NullString
		extra               sql.NullString
		referencedTableName sql.NullString
		constraintTypes     sql.NullString
	)
	for rows.Next() {
		if err := rows.Scan(&tableName, &columnName, &columnDefault, &isNullable, &columnType, &extra, &referencedTableName, &constraintTypes); err != nil {
			return nil, err
		}

		result = append(result, pipe.ColumnMetadata{
			TableName:           tableName.String,
			ColumnName:          columnName.String,
			ColumnDefault:       columnDefault.String,
			IsNullable:          isNullable.String,
			ColumnType:          columnType.String,
			Extra:               extra.String,
			ReferencedTableName: referencedTableName.String,
			ConstraintTypes:     constraintTypes.String,
		})
	}

	return result, nil
}

// columnsがテーブル順であることが前提
func ConvertColumnMetadataToTableMetaData(metadata []pipe.ColumnMetadata) []pipe.TableMetaData {
	result := make([]pipe.TableMetaData, 0, 100)
	currentTableName := ""
	currentColumns := make([]pipe.Column, 0, 20)
	for i, col := range metadata {
		if currentTableName != col.TableName {
			if currentTableName != "" {
				result = append(result, pipe.TableMetaData{TableName: currentTableName, Columns: currentColumns})
			}
			currentTableName = col.TableName
			currentColumns = make([]pipe.Column, 0, 50)
		}
		currentColumns = append(currentColumns, pipe.Column{
			ColumnName:          col.ColumnName,
			ColumnDefault:       col.ColumnDefault,
			IsNullable:          col.IsNullable,
			ColumnType:          col.ColumnType,
			Extra:               col.Extra,
			ReferencedTableName: col.ReferencedTableName,
			ConstraintTypes:     col.ConstraintTypes,
		})

		if i == len(metadata)-1 {
			result = append(result, pipe.TableMetaData{TableName: currentTableName, Columns: currentColumns})
		}

	}
	return result
}
