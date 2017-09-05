package handler

import (
	"strings"

	"github.com/micro/go-log"

	"golang.org/x/net/context"

	proto "github.com/lukyth/try-go-micro/svc_string/proto"
)

// String is a struct
type String struct{}

// Uppercase is a function
func (s *String) Uppercase(ctx context.Context, req *proto.UppercaseRequest, rsp *proto.UppercaseResponse) error {
	log.Log("Received String.Uppercase request")
	rsp.Text = strings.ToUpper(req.Text)
	return nil
}

// Count is a function
func (s *String) Count(ctx context.Context, req *proto.CountRequest, rsp *proto.CountResponse) error {
	log.Log("Received String.Count request")
	rsp.Count = int32(len(req.Text))
	return nil
}
