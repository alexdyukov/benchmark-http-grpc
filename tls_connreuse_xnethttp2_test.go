package main_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
	"golang.org/x/net/http2"
)

func BenchmarkTLSConnReuseXNETHTTP2(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"tls"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		httpClient := &http.Client{
			Transport: &http2.Transport{
				DisableCompression: true,
				ConnPool:           nil, // enable connection reuse
				TLSClientConfig:    globalTLSConfig,
			},
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://localhost:60005/", nil)
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

			if response.Response != "Hello tls on HTTP/2.0" {
				b.Fatal("invalid return value: " + response.Response)
			}
		}
	})
}
