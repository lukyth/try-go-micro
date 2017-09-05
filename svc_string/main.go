package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/lukyth/try-go-micro/svc_string/handler"
	"github.com/lukyth/try-go-micro/svc_string/subscriber"

	example "github.com/lukyth/try-go-micro/svc_string/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.svc_string"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.svc_string", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.svc_string", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
