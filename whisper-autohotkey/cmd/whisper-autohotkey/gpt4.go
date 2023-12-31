package main

import (
	"context"
	"math"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func BuildCommand(config Config, prompt string) (string, error) {

	if strings.TrimSpace(prompt) == "" {
		return "MsgBox, 32,, No input detected! Is your microphone working correctly?", nil
	}

	systemContext, err := os.ReadFile("./prompt.txt")
	if err != nil {
		return "", err
	}

	c := openai.NewClient(config.OpenapiKey)

	response, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			// https://github.com/sashabaranov/go-openai#why-dont-we-get-the-same-answer-when-specifying-a-temperature-field-of-0-and-asking-the-same-question
			Temperature: math.SmallestNonzeroFloat32,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: string(systemContext),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "ACTION: " + prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
