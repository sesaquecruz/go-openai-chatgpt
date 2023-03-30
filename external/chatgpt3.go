package external

import (
	"bytes"
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

func (c *ChatGpt3) Talk(ctx context.Context, message string, response chan<- string) {
	defer close(response)

	userMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}

	c.Messages = append(c.Messages, userMessage)

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

	var responseBuffer bytes.Buffer

	for {
		data, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Panicln(err)
		}

		content := data.Choices[0].Delta.Content

		_, err = responseBuffer.WriteString(content)
		if err != nil {
			log.Panicln(err)
		}

		response <- content
	}

	gptMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: responseBuffer.String(),
	}

	c.Messages = append(c.Messages, gptMessage)
}

func (c *ChatGpt3) ClearHistory() {
	c.Messages = make([]openai.ChatCompletionMessage, 0)
}
