module github.com/alexdyukov/benchmark-http-grpc

go 1.26.0

replace github.com/alexdyukov/benchmark-http-grpc => ./

require (
	github.com/goccy/go-json v0.12.0
	github.com/quic-go/quic-go v0.64.1
	golang.org/x/net v0.62.0
	google.golang.org/grpc v1.85.0
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/quic-go/qpack v0.6.0 // indirect
	go.uber.org/mock v0.6.0 // indirect
	golang.org/x/crypto v0.58.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/text v0.43.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20261026165753-2ffab32bce7e // indirect
)
