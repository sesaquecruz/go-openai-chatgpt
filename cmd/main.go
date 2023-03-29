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

func menu() (int, error) {
	fmt.Println("\nSelect an option:")
	fmt.Println("[1] Batch mode")
	fmt.Println("[2] Stream mode")
	fmt.Print("\n[?] ")

	var option int
	_, err := fmt.Scan(&option)
	if err != nil || option < 1 || option > 2 {
		return 0, err
	}

	return option, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("the .env file was not found")
	}

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatalln("the API_KEY was not found in the .env file")
	}

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	chatGpt3 := external.NewChatGpt3(openai.NewClient(apiKey))

	option, err := menu()
	if err != nil {
		log.Fatalln("invalid option")
	}

	if option == 1 {
		internal.StartChatBatch(ctx, reader, writer, chatGpt3)
	} else {
		internal.StartChatStream(ctx, reader, writer, chatGpt3)
	}
}
