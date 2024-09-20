package databases

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

func DatabaseConnectDescribeGet(state *server.ServerState) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		databaseConnect := state.DatabaseConnect

		response, errMarshal := json.Marshal(databaseConnect)

		if errMarshal != nil {
			helpers.ErrorResponse("Error marshalling database connection info", http.StatusBadRequest, res)
			return
		}

		res.Write(response)
	}
}
