package external

import (
	"context"
	"errors"
	"io"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGpt3 struct {
	Client   *openai.Client
	Model    string
	Messages []openai.ChatCompletionMessage
}

func NewChatGpt3(client *openai.Client) *ChatGpt3 {
	return &ChatGpt3{
		Client:   client,
		Model:    openai.GPT3Dot5Turbo0301,
		Messages: make([]openai.ChatCompletionMessage, 0),
	}
}

func (c *ChatGpt3) ClearMessages() {
	c.Messages = make([]openai.ChatCompletionMessage, 0)
}

func (c *ChatGpt3) AddUserMessage(content string) {
	message := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	}

	c.Messages = append(c.Messages, message)
}

func (c *ChatGpt3) AddGptMessage(content string) {
	message := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	}

	c.Messages = append(c.Messages, message)
}

func (c *ChatGpt3) UpdateChat(ctx context.Context, response chan<- string) {
	defer close(response)

	request := openai.ChatCompletionRequest{
		Model:    c.Model,
		Messages: c.Messages,
		Stream:   true,
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
