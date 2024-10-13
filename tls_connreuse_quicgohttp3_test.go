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

func BenchmarkTLSConnReuseQUICGOHTTP3(b *testing.B) {
	ctx := context.Background()

	bodyRaw := []byte(`{"secure":"tls"}`)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		response := HTTPResponse{}

		roundTripper := &http3.RoundTripper{
			DisableCompression: true,
			Dial:               nil, // enable connection reuse
			TLSClientConfig:    globalTLSConfig,
		}
		defer roundTripper.Close()

		httpClient := &http.Client{
			Transport: roundTripper,
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://localhost:60006/", nil)
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
