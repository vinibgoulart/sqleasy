package databases

import (
	"github.com/manifoldco/promptui"
	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/databases"
)

func DatabaseConnectAsk() *databases.DatabaseConnect {
	logger := helpers.LoggerCreate("DatabaseConnectAsk")

	promptDatabaseType := promptui.Select{
		Label: "Select database type",
		Items: []string{"postgres"},
	}
	_, databaseType, databaseTypeErr := promptDatabaseType.Run()
	if databaseTypeErr != nil {
		logger.Error(databaseTypeErr.Error())
		return nil
	}

	promptDatabaseHost := promptui.Prompt{
		Label: "Database host",
	}
	databaseHost, databaseHostErr := promptDatabaseHost.Run()
	if databaseHostErr != nil {
		logger.Error(databaseHostErr.Error())
		return nil
	}

	promptDatabasePort := promptui.Prompt{
		Label: "Database port",
	}
	databasePort, databasePortErr := promptDatabasePort.Run()
	if databasePortErr != nil {
		logger.Error(databasePortErr.Error())
		return nil
	}

	promptDatabaseName := promptui.Prompt{
		Label: "Database name",
	}
	databaseName, databaseNameErr := promptDatabaseName.Run()
	if databaseNameErr != nil {
		logger.Error(databaseNameErr.Error())
		return nil
	}

	promptDatabasePassword := promptui.Prompt{
		Label: "Database password",
		Mask:  '*',
	}
	databasePassword, databasePasswordErr := promptDatabasePassword.Run()
	if databasePasswordErr != nil {
		logger.Error(databasePasswordErr.Error())
		return nil
	}

	return &databases.DatabaseConnect{
		DatabaseType: databaseType,
		Host:         databaseHost,
		Port:         databasePort,
		Username:     databaseName,
		Password:     databasePassword,
	}
}
