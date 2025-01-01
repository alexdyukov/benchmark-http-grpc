# benchmark-http-grpc
Performance benchmarking of various version of http with json encoding decoding versus grpc

## How to
```
# regenerate proto contracts
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpcapi/*.proto
# ensure you have up to date dependencies
go get -u ./...
# generate self signed cert
openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes -keyout example.key -out example.crt -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost"
# run benchmarks
go test -bench=. -benchmem -benchtime=100000x
```

## Results
```
$ go test -bench=. -benchmem -benchtime=100000x
warning: GOPATH set to GOROOT (/home/user/go) has no effect
goos: linux
goarch: amd64
pkg: github.com/alexdyukov/benchmark-http-grpc
cpu: AMD Ryzen 7 8845H w/ Radeon 780M Graphics
BenchmarkGRPCRAWConnReuse-16              100000              9591 ns/op            9759 B/op        192 allocs/op
BenchmarkGRPCRAWNoConnReuse-16            100000            108120 ns/op          183461 B/op        930 allocs/op
BenchmarkGRPCTLSConnReuse-16              100000              8938 ns/op            9853 B/op        195 allocs/op
BenchmarkGRPCTLSNoConnReuse-16            100000            928456 ns/op          361984 B/op       2136 allocs/op
BenchmarkHTTP1RAWConnReuse-16             100000              6054 ns/op            6445 B/op         76 allocs/op
BenchmarkHTTP1RAWNoConnReuse-16           100000             22216 ns/op           19043 B/op        143 allocs/op
BenchmarkHTTP1TLSConnReuse-16             100000              6313 ns/op            6543 B/op         80 allocs/op
BenchmarkHTTP1TLSNoConnReuse-16           100000            895523 ns/op          200309 B/op       1285 allocs/op
BenchmarkHTTP2RAWConnReuse-16             100000             10702 ns/op           12860 B/op         89 allocs/op
BenchmarkHTTP2TLSConnReuse-16             100000             10903 ns/op           13867 B/op         96 allocs/op
BenchmarkHTTP2TLSNoConnReuse-16           100000            933833 ns/op          215772 B/op       1455 allocs/op
BenchmarkHTTP3TLSConnReuse-16             100000             22874 ns/op           22280 B/op        210 allocs/op
PASS
ok      github.com/alexdyukov/benchmark-http-grpc       297.645s
```

## What repo missing
There is no benchmarks for
1. http2 (x/net/http2) insecure (h2) without connection reuse (disabled keepalives)
2. quic (quic-go/http3) insecure with and without connection reuse
3. quic (quic-go/http3) tls without connection reuse (disabled keepalives)

cause of hardcoded transports and/or internal connection pools and/or lack of keepalive option support