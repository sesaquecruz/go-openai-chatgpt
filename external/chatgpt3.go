package external

import (
	"context"
	"errors"
	"io"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGpt3 struct {
	Client *openai.Client
	Model  string
	Role   string
}

func NewChatGpt3(client *openai.Client) *ChatGpt3 {
	return &ChatGpt3{
		Client: client,
		Model:  openai.GPT3Dot5Turbo,
		Role:   openai.ChatMessageRoleUser,
	}
}

func (c *ChatGpt3) TalkBatch(ctx context.Context, message string) (*string, error) {
	request := openai.ChatCompletionRequest{
		Model: c.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    c.Role,
				Content: message,
			},
		},
	}

	response, err := c.Client.CreateChatCompletion(ctx, request)
	if err != nil {
		return nil, err
	}

	return &response.Choices[0].Message.Content, nil
}

func (c *ChatGpt3) TalkStream(ctx context.Context, message string, response chan<- string) {
	defer close(response)

	request := openai.ChatCompletionRequest{
		Model: c.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    c.Role,
				Content: message,
			},
		},
		Stream: true,
	}

	stream, err := c.Client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		log.Panicln(err)
	}
	defer stream.Close()

	for {
		data, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}

			log.Panicln(err)
		}

		response <- data.Choices[0].Delta.Content
	}
}
