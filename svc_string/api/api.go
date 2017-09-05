package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"

	proto "github.com/lukyth/try-go-micro/svc_string/proto"
)

type uppercaseResult struct {
	text string
	err  error
}

type countResult struct {
	count int32
	err   error
}

// Service is a service
type Service struct {
	Client client.Client
}

// Uppercase is function
func (s *Service) Uppercase(ctx context.Context, req *proto.UppercaseRequest, rsp *proto.UppercaseResponse) error {
	log.Print("Received String.Uppercase API request")

	// checkin and checkout date query params
	text := req.Text

	// make reqeusts for uppercase text
	uppercaseResultCh := getUppercaseResult(ctx, s.Client, text)

	// wait on uppercase reply
	uppercaseResultReply := <-uppercaseResultCh
	if err := uppercaseResultReply.err; err != nil {
		return errors.InternalServerError("api.svc_string", err.Error())
	}

	rsp.Text = uppercaseResultReply.text
	return nil
}

// Count is function
func (s *Service) Count(ctx context.Context, req *proto.CountRequest, rsp *proto.CountResponse) error {
	log.Print("Received String.Count API request")

	// checkin and checkout date query params
	text := req.Text

	// make reqeusts for count result
	countResultCh := getCountResult(ctx, s.Client, text)

	// wait on count reply
	countResultReply := <-countResultCh
	if err := countResultReply.err; err != nil {
		return errors.InternalServerError("api.svc_string", err.Error())
	}

	rsp.Count = countResultReply.count
	return nil
}

func getUppercaseResult(ctx context.Context, c client.Client, text string) chan uppercaseResult {
	stringClient := proto.NewStringClient("svc_string", c)
	ch := make(chan uppercaseResult, 1)

	go func() {
		res, err := stringClient.Uppercase(ctx, &proto.UppercaseRequest{
			Text: text,
		})

		ch <- uppercaseResult{res.Text, err}
	}()

	return ch
}

func getCountResult(ctx context.Context, c client.Client, text string) chan countResult {
	stringClient := proto.NewStringClient("svc_string", c)
	ch := make(chan countResult, 1)

	go func() {
		res, err := stringClient.Count(ctx, &proto.CountRequest{
			Text: text,
		})

		ch <- countResult{res.Count, err}
	}()

	return ch
}

func main() {
	// Service Name need to be prefixed with "go.micro.api" here
	service := micro.NewService(micro.Name("go.micro.api.svc_string"))

	service.Init()

	proto.RegisterStringHandler(service.Server(), &Service{service.Client()})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
