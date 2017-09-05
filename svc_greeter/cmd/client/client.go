package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"

	"github.com/lukyth/try-go-micro/pkg/logging"
	proto "github.com/lukyth/try-go-micro/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.WrapClient(logging.LogWrap),
	)

	// Create new greeter client
	greeter := proto.NewGreeterClient("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Kid"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
