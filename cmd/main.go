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
	godotenv.Load()

	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		log.Fatalln("the API_KEY was not found")
	}

	client := openai.NewClient(apiKey)
	chatGpt3 := external.NewChatGpt3(client)

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	internal.StartChat(ctx, chatGpt3, reader, writer)
}
