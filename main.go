package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/sashabaranov/go-openai"
)

type AIProvider interface {
	Ask(question string) string
}

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

type RandomAI struct{}

func (r *RandomAI) Ask(question string) string {
	responses := []string{
		"Yes, definitely!",
		"No, I don't think so.",
		"Maybe, it's hard to say.",
		"Ask again later.",
	}
	return fmt.Sprintf("Random AI replies: %s", responses[rand.Intn(len(responses))])
}

type ChatBot struct {
	client AIProvider
}

func NewChatBot(client AIProvider) *ChatBot {
	return &ChatBot{client: client}
}

func (c *ChatBot) SetClient(client AIProvider) {
	c.client = client
}

func (c *ChatBot) Ask(question string) string {
	return c.client.Ask(question)
}

func main() {
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	openAI := NewOpenAI(openAIAPIKey)
	randomAI := &RandomAI{}

	bot := NewChatBot(openAI)

	fmt.Println(bot.Ask("How much is Product X?"))

	bot.SetClient(randomAI)
	fmt.Println(bot.Ask("How much is Product X?"))
}
