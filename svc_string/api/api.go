package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	web "github.com/micro/go-web"
	"golang.org/x/net/context"

	proto "github.com/lukyth/try-go-micro/svc_string/proto"
	"github.com/micro/go-micro/client"
)

// String is struct
type String struct{}

// Uppercase is a function
func (s *String) Uppercase(req *restful.Request, rsp *restful.Response) {
	log.Print("Received String.Uppercase API request")

	uppercaseRequest := new(proto.UppercaseRequest)
	err := req.ReadEntity(uppercaseRequest)
	if err != nil {
		rsp.AddHeader("Content-Type", "text/plain")
		rsp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	stringClient := proto.NewStringClient("string", client.DefaultClient)
	response, err := stringClient.Uppercase(context.TODO(), &proto.UppercaseRequest{
		Text: uppercaseRequest.Text,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

// Count is a function
func (s *String) Count(req *restful.Request, rsp *restful.Response) {
	log.Print("Received String.Count API request")

	countRequest := new(proto.CountRequest)
	err := req.ReadEntity(countRequest)
	if err != nil {
		rsp.AddHeader("Content-Type", "text/plain")
		rsp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	stringClient := proto.NewStringClient("string", client.DefaultClient)
	response, err := stringClient.Count(context.TODO(), &proto.CountRequest{
		Text: countRequest.Text,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

func main() {
	// Create service
	service := web.NewService(
		web.Name("api.string"),
	)

	service.Init()

	// Create RESTful handler
	stringHandler := new(String)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path("/string")
	ws.Route(ws.POST("/uppercase").To(stringHandler.Uppercase))
	ws.Route(ws.POST("/count").To(stringHandler.Count))
	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}