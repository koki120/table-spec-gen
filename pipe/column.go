package pipe

import (
	"database/sql"
)

type Column struct {
	TableName           string
	ColumnName          string
	ColumnDefault       string
	IsNullable          string
	ColumnType          string
	Extra               string
	ReferencedTableName string
	ConstraintTypes     string
}

func RowsToColumns(rows *sql.Rows) ([]Column, error) {
	defer rows.Close()
	result := make([]Column, 0, 1000)
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

		result = append(result, Column{
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
