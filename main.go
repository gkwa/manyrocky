package main

import (
	"fmt"
	"os"
)

func main() {
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	openAI := NewOpenAI(openAIAPIKey)
	randomAI := &RandomAI{}

	bot := NewChatBot(openAI)

	fmt.Println(bot.Ask("How much is Product X?"))

	bot.SetClient(randomAI)
	fmt.Println(bot.Ask("How much is Product X?"))
}
