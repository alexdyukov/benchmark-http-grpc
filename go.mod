module github.com/alexdyukov/benchmark-http-grpc

go 1.26.0

replace github.com/alexdyukov/benchmark-http-grpc => ./

require (
	github.com/goccy/go-json v0.10.6
	github.com/quic-go/quic-go v0.62.0
	golang.org/x/net v0.58.0
	google.golang.org/grpc v1.83.2
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/quic-go/qpack v0.6.0 // indirect
	go.uber.org/mock v0.6.0 // indirect
	golang.org/x/crypto v0.55.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260831171406-18b4a7587f8a // indirect
)
