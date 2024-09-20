package databases

import (
	"net/http"

	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/databases"
	"github.com/vinibgoulart/sqleasy/packages/server"
	zius "github.com/vinibgoulart/zius"
)

func DatabaseConnectPost(state *server.ServerState) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var databaseConnect databases.DatabaseConnect

		errJsonDecode := helpers.JsonDecode(res, req, &databaseConnect)
		if errJsonDecode != nil {
			helpers.ErrorResponse(errJsonDecode.Error(), http.StatusBadRequest, res)
			return
		}

		_, errValidate := zius.Validate(databaseConnect)
		if errValidate != nil {
			helpers.ErrorResponse(errValidate.Error(), http.StatusBadRequest, res)
			return
		}

		db, errDatabaseConnect := databases.DatabaseConnectFn(&databaseConnect)

		if errDatabaseConnect != nil {
			helpers.ErrorResponse(errDatabaseConnect.Message, http.StatusBadRequest, res)
			return
		}

		state.Db = db
		state.DatabaseConnect = &databaseConnect

		res.Write([]byte("OK"))
	}
}
