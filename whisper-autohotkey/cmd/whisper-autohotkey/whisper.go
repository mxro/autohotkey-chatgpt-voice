package main

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func Transcribe(inputFileName string, config Config) (string, error) {
	c := openai.NewClient(config.OpenapiKey)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		Prompt:   "",
		FilePath: inputFileName,
	}
	response, err := c.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}

	return response.Text, nil
}
