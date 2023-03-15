package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/raff/tls-ext"
	psk "github.com/raff/tls-psk"
)

type M map[string]interface{}

var (
	Config = &tls.Config{
		ServerName:               "client-PSK-test",
		PreferServerCipherSuites: true,
		CipherSuites:             []uint16{psk.TLS_PSK_WITH_AES_128_CBC_SHA},
		Extra: psk.PSKConfig{
			GetIdentity: func() string {
				return "123"
			},
			GetKey: func(identity string) ([]byte, error) {
				hexString := "166ACC41EC1D4E1DD001ECC130ED0810"
				key := make([]byte, hex.DecodedLen(len(hexString)))
				_, err := hex.Decode(key, []byte(hexString))
				if err != nil {
					return nil, err
				}
				return key, nil
			},
		},
	}
)

func main() {
	baseURL := "https://localhost:9080"
	data := M{"Name": "Noval Agung"}

	responseBody, err := doRequest(baseURL+"/data", "POST", data)
	if err != nil {
		log.Println("ERROR: do request:", err.Error())
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
			return nil, fmt.Errorf("error json encoding: %v", err)
		}
	}

	request, err := http.NewRequest(method, url, payload)
	request.Close = true
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %v", err)
	}

	t := &http.Transport{
		DialTLS: func(network, addr string) (net.Conn, error) {
			return tls.Dial(network, addr, Config)
		},
	}

	client := &http.Client{
		Timeout:   120 * time.Second,
		Transport: t,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error client.Do(): %v", err)
	}
	if response != nil {
		defer response.Body.Close()
	}

	responseBody := make(M)
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
