package databases

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/databases"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

func DatabaseConnectAsk(state *server.ServerState) error {
	logger := helpers.LoggerCreate("Database Connect Ask")
	defer logger.Info("Database connect")

	promptDatabaseType := promptui.Select{
		Label: "Select database type",
		Items: []string{"postgres"},
	}
	_, databaseType, databaseTypeErr := promptDatabaseType.Run()
	if databaseTypeErr != nil {
		logger.Error(databaseTypeErr.Error())
		return databaseTypeErr
	}

	promptDatabaseHost := promptui.Prompt{
		Label:   "Database host",
		Default: "localhost",
	}
	databaseHost, databaseHostErr := promptDatabaseHost.Run()
	if databaseHostErr != nil {
		logger.Error(databaseHostErr.Error())
		return databaseHostErr
	}

	promptDatabasePort := promptui.Prompt{
		Label:   "Database port",
		Default: "5432",
	}
	databasePort, databasePortErr := promptDatabasePort.Run()
	if databasePortErr != nil {
		logger.Error(databasePortErr.Error())
		return databasePortErr
	}

	promptDatabaseName := promptui.Prompt{
		Label:   "Database name",
		Default: "postgres",
	}
	databaseName, databaseNameErr := promptDatabaseName.Run()
	if databaseNameErr != nil {
		logger.Error(databaseNameErr.Error())
		return databaseNameErr
	}

	promptDatabasePassword := promptui.Prompt{
		Label: "Database password",
		Mask:  '*',
	}
	databasePassword, databasePasswordErr := promptDatabasePassword.Run()
	if databasePasswordErr != nil {
		logger.Error(databasePasswordErr.Error())
		return databasePasswordErr
	}

	promptDatabaseUsername := promptui.Prompt{
		Label:   "Database username",
		Default: "postgres",
	}
	databaseUsername, databaseUsernameErr := promptDatabaseUsername.Run()
	if databaseUsernameErr != nil {
		logger.Error(databaseUsernameErr.Error())
		return databaseUsernameErr
	}

	databaseConnect := databases.DatabaseConnect{
		DatabaseType: databaseType,
		Host:         databaseHost,
		Port:         databasePort,
		Username:     databaseUsername,
		Password:     databasePassword,
		Database:     databaseName,
	}

	db, errDatabaseConnect := databases.DatabaseConnectFn(&databaseConnect)

	if errDatabaseConnect != nil {
		logger.Error(errDatabaseConnect.Message)
		return errors.New(errDatabaseConnect.Message)
	}

	state.Db = db
	state.DatabaseConnect = &databaseConnect

	return nil
}
