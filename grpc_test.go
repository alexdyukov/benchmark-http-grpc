package main_test

import (
	"context"
	"log"
	"net"

	"github.com/alexdyukov/benchmark-http-grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcServer struct {
	grpcapi.UnimplementedGrpcServiceServer
}

func (s *GrpcServer) Hello(_ context.Context, req *grpcapi.GrpcInputName) (*grpcapi.GrpcResponse, error) {
	return &grpcapi.GrpcResponse{
		Response: "Hello " + req.InputName,
	}, nil
}

func rawGRPC() {
	lis, err := net.Listen("tcp", ":60000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	grpcapi.RegisterGrpcServiceServer(srv, &GrpcServer{})

	go func() {
		_ = srv.Serve(lis)
	}()
}

func tlsGRPC() {
	lis, err := net.Listen("tcp", ":60001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("example.crt", "example.key")
	if err != nil {
		log.Fatalf("failed to create grpc transport credentials: %v", err)
	}

	srv := grpc.NewServer(grpc.Creds(creds))

	grpcapi.RegisterGrpcServiceServer(srv, &GrpcServer{})

	go func() {
		_ = srv.Serve(lis)
	}()
}
