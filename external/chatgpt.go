package external

import "context"

type ChatGpt interface {
	Talk(ctx context.Context, message string, response chan<- string)
	ClearHistory()
}
