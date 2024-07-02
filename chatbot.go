package main

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
