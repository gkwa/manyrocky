package main

type AIProvider interface {
	Ask(question string) string
}
