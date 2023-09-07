package main

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

func Transcribe(config Config) (string, error) {
	c := openai.NewClient(config.OpenapiKey)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		Prompt:   "Shalveena",
		FilePath: "rec.mp3",
	}
	response, err := c.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}
	e := os.Remove("rec.mp3")
	if e != nil {
		return "", err
	}
	return response.Text, nil
}
