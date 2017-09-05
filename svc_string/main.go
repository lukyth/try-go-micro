package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/lukyth/try-go-micro/svc_string/handler"
	proto "github.com/lukyth/try-go-micro/svc_string/proto"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("svc_string"),
		micro.Version("1.0.0"),
	)

	// Register Handler
	proto.RegisterStringHandler(service.Server(), new(handler.String))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
