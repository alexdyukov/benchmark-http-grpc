package main_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
)

func BenchmarkHTTP1RAWConnReuse(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"raw"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		transport := &http.Transport{
			DisableCompression: true,
			DisableKeepAlives:  false,
			ForceAttemptHTTP2:  false,
		}

		httpClient := &http.Client{Transport: transport}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:40000/", nil)
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
