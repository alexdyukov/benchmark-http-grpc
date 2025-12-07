module github.com/alexdyukov/benchmark-http-grpc

go 1.24.0

toolchain go1.24.2

replace github.com/alexdyukov/benchmark-http-grpc => ./

require (
	github.com/goccy/go-json v0.10.3
	github.com/quic-go/quic-go v0.47.0
	golang.org/x/net v0.47.0
	google.golang.org/grpc v1.77.0
	google.golang.org/protobuf v1.36.10
)

require (
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/pprof v0.0.0-20241017200806-017d972448fc // indirect
	github.com/onsi/ginkgo/v2 v2.20.2 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	go.uber.org/mock v0.5.0 // indirect
	golang.org/x/crypto v0.44.0 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
)
