package main_test

import (
	"context"
	"testing"

	"github.com/alexdyukov/benchmark-http-grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func BenchmarkGRPCTLSConnReuse(b *testing.B) {
	ctx := context.Background()

	creds, err := credentials.NewClientTLSFromFile("example.crt", "localhost")
	if err != nil {
		b.Fatal(err)
	}

	authOption := grpc.WithTransportCredentials(creds)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		conn, err := grpc.NewClient("localhost:30001", authOption)
		if err != nil {
			b.Fatal(err)
		}

		client := grpcapi.NewGrpcServiceClient(conn)

		for pb.Next() {
			resp, err := client.Hello(ctx, &grpcapi.GrpcInputName{InputName: "grpc"})
			if err != nil {
				b.Fatal(err)
			}

			if resp.Response != "Hello grpc" {
				b.Fatal("invalid return value")
			}
		}
	})
}
