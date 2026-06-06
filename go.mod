module github.com/alexdyukov/benchmark-http-grpc

go 1.25.0

replace github.com/alexdyukov/benchmark-http-grpc => ./

require (
	github.com/goccy/go-json v0.10.6
	github.com/quic-go/quic-go v0.60.0
	golang.org/x/net v0.55.0
	google.golang.org/grpc v1.81.1
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/quic-go/qpack v0.6.0 // indirect
	go.uber.org/mock v0.6.0 // indirect
	golang.org/x/crypto v0.52.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
)
