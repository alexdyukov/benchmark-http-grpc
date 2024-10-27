package main_test

import (
	"context"
	"testing"

	"github.com/alexdyukov/benchmark-http-grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BenchmarkGRPCRAWNoConnReuse(b *testing.B) {
	ctx := context.Background()
	authOption := grpc.WithTransportCredentials(insecure.NewCredentials())

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		conn := grpcNoConnReuseClient{
			target:     "localhost:30000",
			dialOption: authOption,
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
