package cli

import (
	"context"
	"sync"

	"github.com/vinibgoulart/sqleasy/helpers"
	aiHandlers "github.com/vinibgoulart/sqleasy/packages/ai/handlers/cli"
	databasesHandlers "github.com/vinibgoulart/sqleasy/packages/databases/handlers/cli"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

var state server.ServerState

func Exec(ctx context.Context, waitGroup *sync.WaitGroup) {
	logger := helpers.LoggerCreate("CLI Exec")
	defer waitGroup.Done()
	defer logger.Info("CLI finished")

	errDatabaseConnect := databasesHandlers.DatabaseConnectAsk(&state)
	if errDatabaseConnect != nil {
		return
	}

	for {
		erroAiPrompt := aiHandlers.AiPromptAsk(&state)
		if erroAiPrompt != nil {
			break
		}
	}
}
