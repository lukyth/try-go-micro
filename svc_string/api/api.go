package main

import (
	"encoding/json"
	"log"

	proto "github.com/lukyth/try-go-micro/svc_string/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"

	"golang.org/x/net/context"
)

// String is struct
type String struct {
	Client proto.StringClient
}

// Uppercase is function
func (s *String) Uppercase(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received String.Uppercase API request")

	response, err := s.Client.Uppercase(ctx, &proto.UppercaseRequest{
		Text: req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"text": response.Text,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	// Service Name need to be prefixed with "go.micro.api" here
	service := micro.NewService(micro.Name("go.micro.api.svc_string"))

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&String{Client: proto.NewStringClient("svc_string", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
