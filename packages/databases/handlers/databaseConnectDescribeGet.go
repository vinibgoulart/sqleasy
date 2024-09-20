package databases

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/databases"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

func DatabaseConnectDescribeGet(state *server.ServerState) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		databaseConnect := state.DatabaseConnect
		db := state.Db

		databaseDescribe, errDatabaseDescribe := databases.DatabaseDescribeFn(databaseConnect, db)

		if errDatabaseDescribe != nil {
			helpers.ErrorResponse(errDatabaseDescribe.Message, http.StatusBadRequest, res)
			return
		}

		response, errMarshal := json.Marshal(databaseDescribe)

		if errMarshal != nil {
			helpers.ErrorResponse("Error marshalling database describe info", http.StatusBadRequest, res)
			return
		}

		res.Write(response)
	}
}
