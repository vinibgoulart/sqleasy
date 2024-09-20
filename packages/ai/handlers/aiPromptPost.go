package ai

import (
	"net/http"

	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/ai"
	"github.com/vinibgoulart/sqleasy/packages/server"
	zius "github.com/vinibgoulart/zius"
)

func AiPromptPost(state *server.ServerState) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var aiPrompt ai.AiPrompt

		errJsonDecode := helpers.JsonDecode(res, req, &aiPrompt)
		if errJsonDecode != nil {
			helpers.ErrorResponse(errJsonDecode.Error(), http.StatusBadRequest, res)
			return
		}

		_, errValidate := zius.Validate(aiPrompt)
		if errValidate != nil {
			helpers.ErrorResponse(errValidate.Error(), http.StatusBadRequest, res)
			return
		}

		resAiPrompt, errAiPrompt := ai.AiPromptGpt(*state, aiPrompt.Prompt)

		if errAiPrompt != nil {
			helpers.ErrorResponse(errAiPrompt.Message, http.StatusBadRequest, res)
			return
		}

		res.Write([]byte(resAiPrompt))
	}
}
