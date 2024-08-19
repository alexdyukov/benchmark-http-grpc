package main_test

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/goccy/go-json"
	"github.com/quic-go/quic-go/http3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type HTTPRequest struct {
	Secure string `json:"secure"`
}

type HTTPResponse struct {
	Response string `json:"response"`
}

func HTTPHello(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := HTTPRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(HTTPResponse{Response: "Hello " + req.Secure + " on " + r.Proto}); err != nil {
		panic(err)
	}
}

func rawNetHTTP1() {
	handler := http.HandlerFunc(HTTPHello)

	server := http.Server{
		Addr:         ":60002",
		Handler:      handler,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){}, // disable http2
	}

	go func() {
		_ = server.ListenAndServe()
	}()
}

func tlsNetHTTP1() {
	handler := http.HandlerFunc(HTTPHello)

	server := http.Server{
		Addr:         ":60003",
		Handler:      handler,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){}, // disable http2
	}

	go func() {
		_ = server.ListenAndServeTLS("example.crt", "example.key")
	}()
}

func rawXNetHTTP2() {
	handler := http.HandlerFunc(HTTPHello)

	h2s := &http2.Server{}
	server := &http.Server{
		Addr:         ":60004",
		Handler:      h2c.NewHandler(handler, h2s),
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	if err := http2.ConfigureServer(server, h2s); err != nil {
		panic(err)
	}

	go func() {
		_ = server.ListenAndServe()
	}()
}

func tlsXNetHTTP2() {
	handler := http.HandlerFunc(HTTPHello)

	server := &http.Server{
		Addr:         ":60005",
		Handler:      handler,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		_ = server.ListenAndServeTLS("example.crt", "example.key")
	}()
}

func tlsQUICGOHTTP3() {
	handler := http.HandlerFunc(HTTPHello)

	server := http3.Server{
		Addr:        ":60006",
		Handler:     handler,
		IdleTimeout: time.Second,
	}

	go func() {
		_ = server.ListenAndServeTLS("example.crt", "example.key")
	}()
}