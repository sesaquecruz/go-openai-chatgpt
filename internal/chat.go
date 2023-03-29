package internal

import (
	"bufio"
	"context"
	"fmt"
	"log"

	"github.com/sesaquecruz/go-openai-chatgpt/external"
)

func StartChatBatch(ctx context.Context, reader *bufio.Reader, writer *bufio.Writer, chatGpt external.ChatGpt) {
	writer.WriteString("\nChatGpt (batch response mode)\n\n")
	if err := writer.Flush(); err != nil {
		log.Panicln(err)
	}

	for {
		writer.WriteString("\n> ")
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Panicln(err)
		}

		response, err := chatGpt.TalkBatch(ctx, message)
		if err != nil {
			log.Panicln(err)
		}

		writer.WriteString(fmt.Sprintf("\n\n%s\n\n", *response))
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}
	}
}

func StartChatStream(ctx context.Context, reader *bufio.Reader, writer *bufio.Writer, chatGpt external.ChatGpt) {
	writer.WriteString("\nChatGpt (stream response mode)\n\n")
	if err := writer.Flush(); err != nil {
		log.Panicln(err)
	}

	for {
		writer.WriteString("\n> ")
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Panicln(err)
		}

		writer.WriteString("\n\n")
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}

		response := make(chan string)
		go chatGpt.TalkStream(ctx, message, response)

		for content := range response {
			writer.WriteString(content)
			writer.Flush()
		}

		writer.WriteString("\n\n")
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}
	}
}
