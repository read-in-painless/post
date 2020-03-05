package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/read-in-painless/post/handler"
	"github.com/read-in-painless/post/subscriber"

	post "github.com/read-in-painless/post/proto/post"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.post"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	post.RegisterPostHandler(service.Server(), new(handler.Post))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.post", service.Server(), new(subscriber.Post))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
