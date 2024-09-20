package databases

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/vinibgoulart/sqleasy/helpers"
)

var databaseDescribeFunctions = map[string]func(*sql.DB) (*sql.DB, *helpers.Error){
	Postgres: databaseDescribePostgres,
}

func DatabaseDescribeFn(databaseConnect *DatabaseConnect, db *sql.DB) (*sql.DB, *helpers.Error) {
	if connectFn, exists := databaseDescribeFunctions[databaseConnect.DatabaseType]; exists {
		return connectFn(db)
	}

	return nil, helpers.ErrorCreate("Invalid database type")
}

func databaseDescribePostgres(db *sql.DB) (*sql.DB, *helpers.Error) {
	// if err != nil {
	// 	return nil, helpers.ErrorCreate(err.Error())
	// }

	return db, nil
}
