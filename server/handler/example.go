package handler

import (
	"log"

	example "github.com/micro/examples/server/proto/example"
	"github.com/micro/go-micro/server"

	"context"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Println("Received message: ", req.GetName())
	rsp.Msg = server.DefaultOptions().Id + ": Hello " + req.Name
	return nil
}
