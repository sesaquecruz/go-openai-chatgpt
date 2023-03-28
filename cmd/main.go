package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sesaquecruz/go-openai-chatgpt/external"
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

	chatGpt := external.NewChatGpt(apiKey)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		response, err := chatGpt.Talk(context.Background(), message)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("\n\n%s\n\n", *response)
	}
}
