package main_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
	"golang.org/x/net/http2"
)

func BenchmarkRAWConnReuseXNETHTTP2(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"raw"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		httpClient := &http.Client{
			Transport: &http2.Transport{
				DisableCompression: true,
				ConnPool:           nil, // enable connection reuse
				AllowHTTP:          true,
				DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
					return new(net.Dialer).DialContext(ctx, network, addr)
				},
			},
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:60004/", nil)
		if err != nil {
			panic(err)
		}

		for pb.Next() {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyRaw))

			resp, err := httpClient.Do(req)
			if err != nil {
				b.Fatal(err)
			}

			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				fmt.Println()
				b.Fatal(err)
			}

			if response.Response != "Hello raw on HTTP/2.0" {
				b.Fatal("invalid return value: " + response.Response)
			}
		}
	})
}
