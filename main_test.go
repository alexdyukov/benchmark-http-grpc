package main_test

import (
	"crypto/tls"
	"os"
	"testing"
	"time"
)

var globalTLSConfig *tls.Config

func TestMain(m *testing.M) {
	cert, err := tls.LoadX509KeyPair("example.crt", "example.key")
	if err != nil {
		panic(err)
	}

	globalTLSConfig = &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cert}}

	rawGRPC()
	tlsGRPC()

	rawNetHTTP1()
	tlsNetHTTP1()

	rawNetHTTP2()
	tlsNetHTTP2()

	// quic doesnt support raw (no tls) server
	tlsQUICGOHTTP3()

	time.Sleep(time.Second)
	os.Exit(m.Run())
}
