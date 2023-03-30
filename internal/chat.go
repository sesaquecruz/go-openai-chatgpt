package internal

import (
	"bufio"
	"context"
	"log"

	"github.com/sesaquecruz/go-openai-chatgpt/external"
)

func StartChat(ctx context.Context, reader *bufio.Reader, writer *bufio.Writer, chatGpt external.ChatGpt) {
	writer.WriteString("\n[Press 'ctrl + c' to exit]\n")
	if err := writer.Flush(); err != nil {
		log.Panicln(err)
	}

	writer.WriteString("\n  _______        __  ________  __________\n")
	writer.WriteString(" / ___/ /  ___ _/ /_/ ___/ _ \\/_  __/_  /\n")
	writer.WriteString("/ /__/ _ \\/ _ `/ __/ (_ / ___/ / / _/_ <\n")
	writer.WriteString("\\___/_//_/\\_,_/\\__/\\___/_/    /_/ /____/\n\n")
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

		message = message[:len(message)-1]
		response := make(chan string, 2)

		go chatGpt.Talk(ctx, message, response)

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
