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

func BenchmarkHTTP2RAWConnReuse(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"raw"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		dialer := &net.Dialer{}

		transport := &http2.Transport{
			DisableCompression: true,
			AllowHTTP:          true,
			DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
				return dialer.DialContext(ctx, network, addr)
			},
		}

		httpClient := &http.Client{Transport: transport}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:50000/", nil)
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
