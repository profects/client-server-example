package main

import (
	"context"
	"log"
	"strconv"
	"time"

	example "github.com/micro/examples/server/proto/example"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"

	br "github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry/consul"
	tp "github.com/micro/go-plugins/transport/nats"
)

func call(i int, c client.Client) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := c.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John -" + strconv.Itoa(i),
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &example.Response{}

	// Call service
	if err := c.Call(ctx, req, rsp); err != nil {
		log.Println("call error ", i, ": err: ", err, rsp)
		return
	}

	log.Println("Call:", i, "rsp:", rsp.Msg)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.client"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(consul.NewRegistry()),
		micro.Transport(tp.NewTransport()),
		micro.Broker(br.NewBroker()),
		//micro.Client(client.NewClient(client.RequestTimeout(2*time.Second))),
	)
	service.Init()

	go func() {
		time.Sleep(5 * time.Second)
		log.Println("\n--- Call example ---")
		for i := 0; i < 100; i++ {
			call(i, service.Client())
		}
		log.Print("DONE")
	}()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
