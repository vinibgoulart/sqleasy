package databases

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/vinibgoulart/sqleasy/helpers"
)

type DatabaseDescribe struct {
	TableName  string `json:"tableName"`
	ColumnName string `json:"columnName"`
	DataType   string `json:"dataType"`
}

type DatabaseColumnProperties struct {
	ColumnName     string `json:"columnName"`
	ColumnDataType string `json:"columnDataType"`
}

type DatabaseDescribed struct {
	TableName    string                     `json:"tableName"`
	TableColumns []DatabaseColumnProperties `json:"tableColumns"`
}

var databaseDescribeFunctions = map[string]func(*sql.DB) ([]*DatabaseDescribed, *helpers.Error){
	Postgres: databaseDescribePostgres,
}

func DatabaseDescribeFn(databaseConnect *DatabaseConnect, db *sql.DB) ([]*DatabaseDescribed, *helpers.Error) {
	if connectFn, exists := databaseDescribeFunctions[databaseConnect.DatabaseType]; exists {
		return connectFn(db)
	}

	return nil, helpers.ErrorCreate("Invalid database type")
}
func databaseDescribePostgres(db *sql.DB) ([]*DatabaseDescribed, *helpers.Error) {
	query := `
		SELECT table_name, column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = 'public';
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, helpers.ErrorCreate(err.Error())
	}
	defer rows.Close()

	tableColumnsMap := make(map[string][]DatabaseColumnProperties)

	tableNames := make(map[string]string)

	for rows.Next() {
		var tableName, columnName, dataType string
		if err := rows.Scan(&tableName, &columnName, &dataType); err != nil {
			return nil, helpers.ErrorCreate(err.Error())
		}

		tableNames[tableName] = tableName

		tableColumnsMap[tableName] = append(tableColumnsMap[tableName], DatabaseColumnProperties{
			ColumnName:     columnName,
			ColumnDataType: dataType,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, helpers.ErrorCreate(err.Error())
	}

	var databaseDescribed []*DatabaseDescribed
	for tableName, columns := range tableColumnsMap {
		databaseDescribed = append(databaseDescribed, &DatabaseDescribed{
			TableName:    tableName,
			TableColumns: columns,
		})
	}

	return databaseDescribed, nil
}
