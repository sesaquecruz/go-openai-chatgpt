package external

import "context"

type ChatGpt interface {
	ClearMessages()
	AddUserMessage(content string)
	AddGptMessage(content string)
	UpdateChat(ctx context.Context, response chan<- string)
}
