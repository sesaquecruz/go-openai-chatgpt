package external

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGpt struct {
	client *openai.Client
}

func NewChatGpt(apiKey string) *ChatGpt {
	return &ChatGpt{client: openai.NewClient(apiKey)}
}

func (c *ChatGpt) Talk(ctx context.Context, message string) (response *string, err error) {
	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	text := &resp.Choices[0].Message.Content
	return text, nil
}
