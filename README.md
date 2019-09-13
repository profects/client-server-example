# client-server-example

This is a microservice example. Run the server first and the client after. Feel free to bump the go-micro version or something else and check if the microservice keeps running.

## Contents

- server/main.go - initialises and runs the the server
- server/handler - is an example RPC request handler for the Server
- server/proto - contains the protobuf defintion for the Server API

- client/main.go - initialises and runs the the client

## Usage

Run Service

```bash
$ go run server/main.go
2019/09/13 10:52:26 Transport [nats] Listening on _INBOX.7YQlFtWCowlKWyThsLY43S
2019/09/13 10:52:26 Broker [nats] Connected to nats://127.0.0.1:4222
2019/09/13 10:52:26 Registry [consul] Registering node: go.micro.srv.example-469001d1-ff2a-4d40-92b2-9e3c6e7361c7
```

Test Service

```bash
$ go run client/main.go
2019/09/13 10:52:32 Transport [nats] Listening on _INBOX.oblwlIaiO8eTdxjz6GwlY1
2019/09/13 10:52:32 Broker [nats] Connected to nats://127.0.0.1:4222
2019/09/13 10:52:32 Registry [consul] Registering node: go.micro.client-85cce84b-389e-42eb-8258-ceaf02f27158
```

## consul

[Consul with client/server](https://prnt.sc/p5kxt1)
