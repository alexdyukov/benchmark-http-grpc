# benchmark-http-grpc
Performance benchmarking of various version of http with json encoding decoding versus grpc

## How to
```
# regenerate proto contracts
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpcapi/*.proto
# generate self signed cert
openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes -keyout example.key -out example.crt -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost"
# remember of udp buffer sizes, more info at https://github.com/quic-go/quic-go/wiki/UDP-Buffer-Sizes
sudo sysctl -w net.core.rmem_max=7500000
sudo sysctl -w net.core.wmem_max=7500000
# run benchmarks
go test -bench=. -benchmem -benchtime=100000x
```

## Results
```
$ go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: github.com/alexdyukov/benchmark-http-grpc
cpu: AMD Ryzen 7 8845H w/ Radeon 780M Graphics
BenchmarkGRPCRAWConnReuse-16              100000             88381 ns/op            9581 B/op        152 allocs/op
BenchmarkGRPCRAWNoConnReuse-16            100000            434862 ns/op           70205 B/op        877 allocs/op
BenchmarkGRPCTLSConnReuse-16              100000             82215 ns/op            9178 B/op        154 allocs/op
BenchmarkGRPCTLSNoConnReuse-16            100000           1643176 ns/op          384387 B/op       1869 allocs/op
BenchmarkHTTP1RAWConnReuse-16             100000             68501 ns/op            6998 B/op         82 allocs/op
BenchmarkHTTP1RAWNoConnReuse-16           100000            150305 ns/op           20488 B/op        155 allocs/op
BenchmarkHTTP1TLSConnReuse-16             100000             64722 ns/op            7149 B/op         86 allocs/op
BenchmarkHTTP1TLSNoConnReuse-16           100000           1429114 ns/op          278701 B/op       1106 allocs/op
BenchmarkHTTP2RAWConnReuse-16             100000             78032 ns/op           11280 B/op         91 allocs/op
BenchmarkHTTP2TLSConnReuse-16             100000             80200 ns/op           12413 B/op         96 allocs/op
BenchmarkHTTP2TLSNoConnReuse-16           100000           1477753 ns/op          282886 B/op       1274 allocs/op
BenchmarkHTTP3TLSConnReuse-16             100000             78790 ns/op           22781 B/op        220 allocs/op
PASS
ok      github.com/alexdyukov/benchmark-http-grpc       588.957s
```

## What repo missing
There is no benchmarks for
1. http2 (x/net/http2) insecure (h2) without connection reuse (disabled keepalives)
2. quic (quic-go/http3) insecure with and without connection reuse
3. quic (quic-go/http3) tls without connection reuse (disabled keepalives)

cause of hardcoded transports and/or internal connection pools and/or lack of keepalive option support