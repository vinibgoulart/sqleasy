package cli

import (
	"context"
	"fmt"
	"sync"

	"github.com/vinibgoulart/sqleasy/helpers"
	databasesHandlers "github.com/vinibgoulart/sqleasy/packages/databases/handlers/cli"
)

func Exec(ctx context.Context, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	logger := helpers.LoggerCreate("CLI Exec")

	databaseType := databasesHandlers.DatabaseConnectAsk()
	if databaseType == nil {
		logger.Error("Error on database connect")
		return
	}

	fmt.Println(databaseType)
}
