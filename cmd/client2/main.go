package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Load the client's TLS certificate and key
	cert, err := tls.LoadX509KeyPair("../../client.crt", "../../client.key")
	if err != nil {
		log.Fatalf("Failed to load client certificate and key: %v", err)
	}

	// Load the server's self-signed certificate
	serverCert, err := ioutil.ReadFile("../../server.crt")
	if err != nil {
		log.Fatalf("Failed to load server certificate: %v", err)
	}

	// Create a certificate pool containing the server's self-signed certificate
	serverCertPool := x509.NewCertPool()
	serverCertPool.AppendCertsFromPEM(serverCert)

	// Create a TLS configuration that uses the client's certificate and key and verifies the server's self-signed certificate
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      serverCertPool,
	}

	// Create an HTTPS client with the TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Send an HTTPS request to the server

	// ---------- cara 1
	// resp, err := client.Get("https://localhost:9080")

	// --------- cara 2
	// req, err := http.NewRequest("GET", "https://localhost:9080", nil)
	// if err != nil {
	// 	log.Fatalf("Failed to create HTTP request: %v", err)
	// }

	// // Explicitly set the server name in the request
	// req.Host = "localhost"

	// resp, err := client.Do(req)

	// -------- cara 3
	req, err := http.NewRequest("GET", "https://127.0.0.1:9080", nil)
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed to send HTTPS request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Print the response body
	log.Printf("Response body: %s", body)
}
