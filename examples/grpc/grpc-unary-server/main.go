package main

import (
	"gofr.dev/examples/grpc/grpc-unary-server/server"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New(nil)

	server.RegisterHelloServerWithGofr(app, server.NewHelloGoFrServer())

	app.Run()
}
