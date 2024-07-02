package main

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openai.Client
}

func NewOpenAI(apiKey string) *OpenAI {
	return &OpenAI{
		client: openai.NewClient(apiKey),
	}
}

func (o *OpenAI) Ask(question string) string {
	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)
	if err != nil {
		return "Error: " + err.Error()
	}

	return "OpenAI says: " + resp.Choices[0].Message.Content
}
