package main

import (
	"log"
	"time"

	"git.profects.com/profects/server/server/handler"
	micro "github.com/micro/go-micro"
	br "github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	tp "github.com/micro/go-plugins/transport/nats"
)

func main() {
	// Initialise Server
	srv := server.NewServer(
		server.Name("go.micro.srv.example"),
	)

	service := micro.NewService(
		micro.Server(srv),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(consul.NewRegistry()),
		micro.Transport(tp.NewTransport()),
		micro.Broker(br.NewBroker()),
	)
	service.Init()

	err := srv.Handle(
		server.NewHandler(new(handler.Example)),
	)
	if err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
