package logging

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/server"
)

// LogHandlerWrapper implements the server.HandlerWrapper
func LogHandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[Log Wrapper] Before serving request method: %v\n", req.Method())
		err := fn(ctx, req, rsp)
		fmt.Println("[Log Wrapper] After serving request")
		return err
	}
}
