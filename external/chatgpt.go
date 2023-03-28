package external

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGpt struct {
	Client *openai.Client
	Model  string
	Role   string
}

func NewChatGpt(client *openai.Client, model string, role string) *ChatGpt {
	return &ChatGpt{
		Client: client,
		Model:  model,
		Role:   role,
	}
}

func (c *ChatGpt) TalkBatchResponse(ctx context.Context, message string) (*string, error) {
	response, err := c.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: c.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    c.Role,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	content := &response.Choices[0].Message.Content
	return content, nil
}
