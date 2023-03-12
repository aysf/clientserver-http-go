package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	tls "github.com/raff/tls-ext"
	psk "github.com/raff/tls-psk"
)

type M map[string]interface{}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/data", ActionData)

	// server := new(http.Server)
	// server.Handler = mux
	// server.Addr = ":9080"
	// log.Println("Starting server at", server.Addr)

	s := &http.Server{
		Addr:    ":9080",
		Handler: mux,
	}

	l, err := net.Listen("tls", ":9080")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}

	cfg := &tls.Config{
		CipherSuites: []uint16{psk.TLS_DHE_PSK_WITH_AES_128_CBC_SHA, psk.TLS_DHE_PSK_WITH_AES_256_CBC_SHA},
		Certificates: []tls.Certificate(tls.Certificate{}),
		MaxVersion:   tls.VersionTLS12,
		Extra: psk.PSKConfig{
			GetKey: getIdentityKey,
		},
	}

	l = tls.NewListener(l, cfg)

	s.Serve(l)

	// err := server.ListenAndServeTLS("./secure/server.crt", "./secure/server.key")

}

func getIdentityKey(id string) ([]byte, error) {

	return nil, nil
}

func ActionData(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request with method", r.Method)

	if r.Method != "POST" {
		fmt.Fprintln(w, "helo get:)")
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
