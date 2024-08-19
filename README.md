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
BenchmarkRAWConnReuseGRPC-16              100000              8922 ns/op            9759 B/op        192 allocs/op
BenchmarkRAWConnReuseNETHTTP1-16          100000              6192 ns/op            6534 B/op         76 allocs/op
BenchmarkRAWConnReuseXNETHTTP2-16         100000              8835 ns/op           13899 B/op         89 allocs/op
BenchmarkRAWNoConnReuseNETHTTP1-16        100000             19583 ns/op           19380 B/op        144 allocs/op
BenchmarkTLSConnReuseGRPC-16              100000              8677 ns/op            9855 B/op        195 allocs/op
BenchmarkTLSConnReuseNETHTTP1-16          100000              5416 ns/op            6587 B/op         80 allocs/op
BenchmarkTLSConnReuseQUICGOHTTP3-16       100000             16872 ns/op           22172 B/op        209 allocs/op
BenchmarkTLSConnReuseXNETHTTP2-16         100000              9281 ns/op           12329 B/op         94 allocs/op
BenchmarkTLSNoConnReuseNETHTTP1-16        100000            761961 ns/op          198551 B/op       1286 allocs/op
PASS
ok      github.com/alexdyukov/benchmark-http-grpc       85.752s
```

## TODO tests without keepalive (no conn reuse) for
1. grpc insecure
2. grpc tls
3. x/net/http2 insecure
4. x/net/http2 tls
5. quic-go/http3 tls