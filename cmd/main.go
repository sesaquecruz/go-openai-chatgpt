package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/sesaquecruz/go-openai-chatgpt/external"
	"github.com/sesaquecruz/go-openai-chatgpt/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("the .env file was not found")
	}

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatal("the API_KEY was not found in the .env file")
	}

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	chatGpt := external.NewChatGpt(
		openai.NewClient(apiKey),
		openai.GPT3Dot5Turbo,
		openai.ChatMessageRoleUser,
	)

	internal.ChatBatchResponse(ctx, reader, writer, chatGpt)
}
