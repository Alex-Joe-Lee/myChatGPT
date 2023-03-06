package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

const (
	OpenAIKey = "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"
)

func main() {
	client := openai.NewClient(OpenAIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(resp)
	fmt.Println(resp.Choices[0].Message.Content)
}
