package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/raff/tls-ext"
	psk "github.com/raff/tls-psk"
)

type M map[string]interface{}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/data", ActionData)

	s := &http.Server{
		Addr:    ":9080",
		Handler: mux,
	}

	cfg := &tls.Config{
		CipherSuites: []uint16{
			psk.TLS_PSK_WITH_AES_128_CBC_SHA,
			psk.TLS_PSK_WITH_AES_256_CBC_SHA,
			psk.TLS_PSK_WITH_3DES_EDE_CBC_SHA,
		},
		PreferServerCipherSuites: true,
		Certificates:             []tls.Certificate{{}},
		MaxVersion:               tls.VersionTLS12,
		MinVersion:               tls.VersionTLS11,
		Extra: psk.PSKConfig{
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

	l, err := tls.Listen("tcp", ":9080", cfg)
	// l, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}
	// l = tls.NewListener(l, cfg)

	log.Println("Starting server at 9080")

	if err := s.Serve(l); err != nil {
		log.Fatalln("error starting server:", err)
	}

}

func ActionData(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request with method", r.Method)

	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	payload := make(M)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := payload["Name"]; !ok {
		http.Error(w, "Payload `Name` is required", http.StatusBadRequest)
		return
	}

	data := M{
		"Message": fmt.Sprintf("Hello %s", payload["Name"]),
		"Status":  true,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
