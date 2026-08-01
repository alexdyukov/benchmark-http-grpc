module github.com/alexdyukov/benchmark-http-grpc

go 1.25.0

replace github.com/alexdyukov/benchmark-http-grpc => ./

require (
	github.com/goccy/go-json v0.10.6
	github.com/quic-go/quic-go v0.61.0
	golang.org/x/net v0.57.0
	google.golang.org/grpc v1.83.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/quic-go/qpack v0.6.0 // indirect
	go.uber.org/mock v0.6.0 // indirect
	golang.org/x/crypto v0.54.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260729162451-8efbd57d26e0 // indirect
)
