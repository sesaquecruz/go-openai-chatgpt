package internal

import (
	"bufio"
	"context"
	"fmt"
	"log"

	"github.com/sesaquecruz/go-openai-chatgpt/external"
)

func ChatBatchResponse(
	ctx context.Context,
	reader *bufio.Reader,
	writer *bufio.Writer,
	chatGpt *external.ChatGpt,
) {
	writer.WriteString("\nChatGPT (batch response)\n\n")
	if err := writer.Flush(); err != nil {
		log.Fatalln(err)
	}

	for {
		writer.WriteString("\n> ")
		if err := writer.Flush(); err != nil {
			log.Fatalln(err)
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		response, err := chatGpt.TalkBatchResponse(ctx, message)
		if err != nil {
			log.Fatalln(err)
		}

		writer.WriteString(fmt.Sprintf("\n\n%s\n\n", *response))
		if err := writer.Flush(); err != nil {
			log.Fatalln(err)
		}
	}
}
