package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/sesaquecruz/go-openai-chatgpt/external"
	"github.com/sesaquecruz/go-openai-chatgpt/internal"
)

func menu() int {
	for {
		fmt.Println("\n[Press 'ctrl + c' to exit]")
		fmt.Println("\nSelect an option:")
		fmt.Println("[1] Stream mode")
		fmt.Println("[2] Batch mode")

		fmt.Print("\n[?] ")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			continue
		}

		if option != 1 && option != 2 {
			continue
		}

		return option
	}
}

func main() {
	godotenv.Load()

	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		log.Fatalln("the API_KEY was not found")
	}

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	chatGpt3 := external.NewChatGpt3(openai.NewClient(apiKey))

	option := menu()
	if option == 1 {
		internal.StartChatStream(ctx, reader, writer, chatGpt3)
	} else {
		internal.StartChatBatch(ctx, reader, writer, chatGpt3)
	}
}
