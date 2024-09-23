package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/zius"
)

const (
	Postgres = "postgres"
	MsSql    = "mssql"
)

type DatabaseConnect struct {
	DatabaseType string `json:"databaseType" zius:"equals={mssql}:databaseType must be postgres or mssql,string"`
	Host         string `json:"host" zius:"required:host is required,string"`
	Port         string `json:"port" zius:"required:port is required,string"`
	Username     string `json:"username" zius:"required:username is required,string"`
	Password     string `json:"password" zius:"string"`
	Database     string `json:"database" zius:"required:database is required,string"`
}

var databaseConnectFunctions = map[string]func(*DatabaseConnect) (*sql.DB, *helpers.Error){
	Postgres: databaseConnectPostgres,
	MsSql:    databaseConnectMsSql,
}

func DatabaseConnectFn(databaseConnect *DatabaseConnect) (*sql.DB, *helpers.Error) {
	_, errValidate := zius.Validate(*databaseConnect)
	if errValidate != nil {
		return nil, helpers.ErrorCreate(errValidate.Error())
	}

	if connectFn, exists := databaseConnectFunctions[databaseConnect.DatabaseType]; exists {
		return connectFn(databaseConnect)
	}

	return nil, helpers.ErrorCreate("Invalid database type")
}

func databaseConnectPostgres(databaseConnect *DatabaseConnect) (*sql.DB, *helpers.Error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", databaseConnect.Host, databaseConnect.Port, databaseConnect.Username, databaseConnect.Password, databaseConnect.Database)

	db, err := sql.Open(Postgres, dataSourceName)
	errPing := db.Ping()
	if errPing != nil {
		return nil, helpers.ErrorCreate(errPing.Error())
	}

	if err != nil {
		return nil, helpers.ErrorCreate(err.Error())
	}

	return db, nil
}

func databaseConnectMsSql(databaseConnect *DatabaseConnect) (*sql.DB, *helpers.Error) {
	dataSourceName := fmt.Sprintf("server=%s;user id=%s;database=%s;password=%s;port=%s", databaseConnect.Host, databaseConnect.Username, databaseConnect.Database, databaseConnect.Password, databaseConnect.Port)

	db, err := sql.Open(MsSql, dataSourceName)
	errPing := db.Ping()
	if errPing != nil {
		return nil, helpers.ErrorCreate(errPing.Error())
	}

	if err != nil {
		return nil, helpers.ErrorCreate(err.Error())
	}

	return db, nil
}
