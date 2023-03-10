package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type M map[string]interface{}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/data", ActionData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":9080"

	log.Println("Starting server at", server.Addr)
	// err := server.ListenAndServe()
	err := server.ListenAndServeTLS("./secure/server.crt", "./secure/server.key")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}
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
