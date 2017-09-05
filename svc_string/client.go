package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"

	proto "github.com/lukyth/try-go-micro/svc_string/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("svc_string.client"))

	// Create new string client
	string := proto.NewStringClient("svc_string", service.Client())

	// Call the string
	rsp, err := string.Count(context.TODO(), &proto.CountRequest{Text: "kid"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Count)
}
