package internal

import (
	"bufio"
	"bytes"
	"context"
	"log"

	"github.com/sesaquecruz/go-openai-chatgpt/external"
)

func StartChat(ctx context.Context, chatGpt external.ChatGpt, reader *bufio.Reader, writer *bufio.Writer) {
	writer.WriteString("\n[Press 'ctrl + c' to exit]")
	writer.WriteString("\n[Enter 'clear' to clear the context]\n")
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
		writer.WriteString("\n > ")
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}

		userMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Panicln(err)
		}

		userMessage = userMessage[:len(userMessage)-1]

		if len(userMessage) == 0 {
			continue
		}

		if userMessage == "clear" {
			chatGpt.ClearMessages()
			continue
		}

		chatGpt.AddUserMessage(userMessage)

		response := make(chan string, 2)
		var gptMessage bytes.Buffer

		go chatGpt.UpdateChat(ctx, response)

		writer.WriteString("\n\n")
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}

		for content := range response {
			gptMessage.WriteString(content)
			writer.WriteString(content)
			writer.Flush()
		}

		writer.WriteString("\n\n")
		if err := writer.Flush(); err != nil {
			log.Panicln(err)
		}

		chatGpt.AddGptMessage(gptMessage.String())
	}
}
