package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/databases"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

func AiPromptGpt(state server.ServerState, prompt string) (string, *helpers.Error) {
	logger := helpers.LoggerCreate("AI Prompt GPT")
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	debugMode := os.Getenv("DEBUG") == "true"

	databaseDescribe, errDatabaseDescribe := databases.DatabaseDescribeFn(state.DatabaseConnect, state.Db)
	if errDatabaseDescribe != nil {
		return "", helpers.ErrorCreate(errDatabaseDescribe.Message)
	}

	var values []interface{}
	for _, ptr := range databaseDescribe {
		values = append(values, *ptr)
	}
	fullPrompt := fmt.Sprintf(
		"Database data: %v. Generate a SQL query to retrieve the data considering the following prompt: %s. Only respond with the SQL query. Use the following database type: %s. Only use tables and columns from the database describe. Return in plain text, without mdx syntax.",
		values,
		prompt,
		state.DatabaseConnect.DatabaseType,
	)

	clientChatResponse, clientChatErr := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini20240718,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fullPrompt,
				},
			},
		},
	)

	if clientChatErr != nil {
		return "", helpers.ErrorCreate(clientChatErr.Error())
	}

	clientChatQuery := clientChatResponse.Choices[0].Message.Content

	if strings.Contains(clientChatQuery, "```sql") {
		clientChatQuery = strings.Replace(clientChatQuery, "```sql", "", -1)
		clientChatQuery = strings.Replace(clientChatQuery, "```", "", -1)
	}

	if debugMode {
		logger.Debug(clientChatQuery)
	}

	rows, err := state.Db.Query(clientChatQuery)
	if err != nil {
		return "", helpers.ErrorCreate(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "", helpers.ErrorCreate(err.Error())
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return "", helpers.ErrorCreate(err.Error())
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			row[col] = values[i]
		}

		results = append(results, row)
	}

	resultsJson, err := json.Marshal(results)
	if err != nil {
		return "", helpers.ErrorCreate(err.Error())
	}

	finalPrompt := fmt.Sprintf(
		"Database query results: %s. Based on the data and the original prompt: %s, generate a summary or response.",
		string(resultsJson),
		prompt,
	)

	finalResponse, finalResponseErr := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: finalPrompt,
				},
			},
		},
	)

	if finalResponseErr != nil {
		return "", helpers.ErrorCreate(finalResponseErr.Error())
	}

	finalContent := finalResponse.Choices[0].Message.Content
	return strings.TrimSpace(finalContent), nil
}
