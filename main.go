package main

import (
	"context"
	"fmt"
	"greeting-card-bot/constants"

	openai "github.com/sashabaranov/go-openai"
)

var styles = map[string]string {
	"brodsky": "Бродского",
}

func main() {
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Привет! Напиши, пожалуйста, поздравление с Днем рождения для Ярославы используя цитату %s", constants.SongQuotes[5].Quote),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}