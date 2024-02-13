package main

import (
	"context"
	"fmt"

	"custom-fizzbuzz/internal/handler"
	"custom-fizzbuzz/internal/server"
)

func main() {
	ctx := context.Background()
	h := &handler.Handler{}
	s := server.NewHttpServer(h, ":4000")
	fmt.Println("Starting http server listening on port 4000")
	defer s.Shutdown(ctx)
	go s.ListenAndServe()

	fmt.Println("Starting gRPC Server listening on port 5000")
	err := server.StartGRPCServer(h, 5000)
	if err != nil {
		fmt.Printf("error starting gRPC server: %#v\n", err)
		return
	}
}
