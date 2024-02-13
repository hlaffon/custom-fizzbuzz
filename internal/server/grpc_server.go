package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"custom-fizzbuzz/internal/handler"
	"custom-fizzbuzz/pkg/model"
	"custom-fizzbuzz/pkg/pb"
)

func StartGRPCServer(handler *handler.Handler, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := newGRPCServer(handler, s)

	pb.RegisterApiServer(s, server)
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

type grpcServer struct {
	pb.ApiServer
	handler *handler.Handler
	server  *grpc.Server
}

func newGRPCServer(handler *handler.Handler, server *grpc.Server) *grpcServer {
	return &grpcServer{handler: handler, server: server}
}

func (s *grpcServer) PrintNumbers(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	input := model.FromProto(req)
	err := input.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	numbers := s.handler.PrintNumber(input)
	return model.ToProto(numbers), nil
}
