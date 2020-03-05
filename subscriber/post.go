package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	post "github.com/read-in-painless/post/proto/post"
)

type Post struct{}

func (e *Post) Handle(ctx context.Context, msg *post.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *post.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
