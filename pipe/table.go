package pipe

type Table struct {
	TableName string
	Columns   []column
}

type column struct {
	ColumnName          string
	ColumnDefault       string
	IsNullable          string
	ColumnType          string
	Extra               string
	ReferencedTableName string
	ConstraintTypes     string
}

// columnsがテーブル順であることが前提
func ColumnsToTables(cols []Column) []Table {
	result := make([]Table, 0, 100)
	currentTableName := ""
	currentColumns := make([]column, 0, 20)
	for i, col := range cols {
		if currentTableName != col.TableName || i == len(cols)-1 {
			if currentTableName != "" {
				result = append(result, Table{TableName: currentTableName, Columns: currentColumns})
			}
			currentTableName = col.TableName
			currentColumns = make([]column, 0, 20)
		}
		currentColumns = append(currentColumns, column{
			ColumnName:          col.ColumnName,
			ColumnDefault:       col.ColumnDefault,
			IsNullable:          col.IsNullable,
			ColumnType:          col.ColumnType,
			Extra:               col.Extra,
			ReferencedTableName: col.ReferencedTableName,
			ConstraintTypes:     col.ConstraintTypes,
		})
	}
	return result
}
