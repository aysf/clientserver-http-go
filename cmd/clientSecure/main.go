package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type M map[string]interface{}

func main() {
	baseURL := "https://localhost:9080"
	method := "POST"
	data := M{"Name": "Noval Agung"}

	responseBody, err := doRequest(baseURL+"/data", method, data)
	if err != nil {
		log.Println("ERROR", err.Error())
		return
	}

	log.Printf("%#v \n", responseBody)
}

func doRequest(url, method string, data interface{}) (interface{}, error) {
	var payload *bytes.Buffer = nil

	if data != nil {
		payload = new(bytes.Buffer)
		err := json.NewEncoder(payload).Encode(data)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	certFile, err := os.ReadFile("./secure/server.crt")
	if err != nil {
		return nil, err
	}

	// Create a certificate pool and add the server's certificate to it
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certFile)

	// Create a TLS configuration with the certificate pool
	config := &tls.Config{
		RootCAs: caCertPool,
	}

	// Create a client with the configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	//
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	responseBody := make(M)
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
