package server

import (
	"database/sql"

	"github.com/vinibgoulart/sqleasy/packages/databases"
)

type ServerState struct {
	Db              *sql.DB
	DatabaseConnect *databases.DatabaseConnect
}
