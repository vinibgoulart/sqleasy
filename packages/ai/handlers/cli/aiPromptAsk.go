package ai

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/ai"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

func AiPromptAsk(state *server.ServerState) error {
	logger := helpers.LoggerCreate("AI Prompt Ask")

	promptAsk := promptui.Prompt{
		Label: "Prompt",
	}
	prompt, promptErr := promptAsk.Run()
	if promptErr != nil {
		logger.Error(promptErr.Error())
		return promptErr
	}

	if prompt == "" {
		logger.Error("Prompt is empty")
		return nil
	}

	resAiPrompt, errAiPrompt := ai.AiPromptGpt(*state, prompt)

	if errAiPrompt != nil {
		logger.Error(errAiPrompt.Message)
		return errors.New(errAiPrompt.Message)
	}

	logger.Info(resAiPrompt)
	return nil
}
