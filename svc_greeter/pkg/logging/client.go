package logging

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[Log Wrapper] Client request service: %s method: %s\n", req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// LogWrap implements client.Wrapper as logWrapper
func LogWrap(c client.Client) client.Client {
	return &logWrapper{c}
}
