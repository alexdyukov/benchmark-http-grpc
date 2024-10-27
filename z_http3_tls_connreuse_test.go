package main_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
	"github.com/quic-go/quic-go/http3"
)

func BenchmarkHTTP3TLSConnReuse(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"tls"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		transport := &http3.RoundTripper{
			DisableCompression: true,
			TLSClientConfig:    globalTLSConfig,
		}
		defer transport.Close()

		httpClient := &http.Client{Transport: transport}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://localhost:60001/", nil)
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

			if response.Response != "Hello tls on HTTP/3.0" {
				b.Fatal("invalid return value: " + response.Response)
			}
		}
	})
}
