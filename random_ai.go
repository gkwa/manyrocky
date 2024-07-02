package main

import (
	"fmt"
	"math/rand"
)

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
