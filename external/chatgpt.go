package external

import "context"

type ChatGpt interface {
	TalkBatch(ctx context.Context, message string) (*string, error)
	TalkStream(ctx context.Context, message string, response chan<- string)
}
