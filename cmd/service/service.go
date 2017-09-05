package main

import (
	"fmt"

	"github.com/micro/go-micro"
	"golang.org/x/net/context"

	"github.com/lukyth/try-go-micro/pkg/logging"
	proto "github.com/lukyth/try-go-micro/proto"
)

// Greeter is a struct
type Greeter struct{}

// Hello is a function
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.0"),
		micro.WrapHandler(logging.LogHandlerWrapper),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
