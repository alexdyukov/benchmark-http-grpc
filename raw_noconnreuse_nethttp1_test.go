package main_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
)

func BenchmarkRAWNoConnReuseNETHTTP1(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"raw"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		httpClient := &http.Client{
			Transport: &http.Transport{
				DisableCompression: true,
				DisableKeepAlives:  true,  // disable connection reuse
				ForceAttemptHTTP2:  false, // http1.1 only
			},
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:60002/", nil)
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
				b.Fatal(err)
			}

			if response.Response != "Hello raw on HTTP/1.1" {
				b.Fatal("invalid return value: " + response.Response)
			}
		}
	})
}
