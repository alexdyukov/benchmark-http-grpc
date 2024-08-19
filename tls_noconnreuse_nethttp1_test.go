package main_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
)

func BenchmarkTLSNoConnReuseNETHTTP1(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"tls"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		httpClient := &http.Client{
			Transport: &http.Transport{
				DisableCompression: true,
				DisableKeepAlives:  true,  // disable connection reuse
				ForceAttemptHTTP2:  false, // http1.1 only
				TLSClientConfig:    globalTLSConfig,
			},
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://localhost:60003/", nil)
		if err != nil {
			b.Fatal(err)
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

			if response.Response != "Hello tls on HTTP/1.1" {
				b.Fatal("invalid return value: " + response.Response)
			}
		}
	})
}
