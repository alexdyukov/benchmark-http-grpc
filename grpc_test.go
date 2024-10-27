package main_test

import (
	"context"
	"log"
	"net"

	"github.com/alexdyukov/benchmark-http-grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type grpcNoConnReuseClient struct {
	target     string
	dialOption grpc.DialOption
}

func (c grpcNoConnReuseClient) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	// new conn every Invoke
	conn, err := grpc.DialContext(ctx, c.target, c.dialOption)
	if err != nil {
		return err
	}

	defer func() {
		go conn.Close()
	}()

	return conn.Invoke(ctx, method, args, reply, opts...)
}

func (grpcNoConnReuseClient) NewStream(_ context.Context, _ *grpc.StreamDesc, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
	// unused in test method
	return nil, nil
}

type GrpcServer struct {
	grpcapi.UnimplementedGrpcServiceServer
}

func (s *GrpcServer) Hello(_ context.Context, req *grpcapi.GrpcInputName) (*grpcapi.GrpcResponse, error) {
	return &grpcapi.GrpcResponse{
		Response: "Hello " + req.InputName,
	}, nil
}

func rawGRPC() {
	lis, err := net.Listen("tcp", ":30000")
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
	lis, err := net.Listen("tcp", ":30001")
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
