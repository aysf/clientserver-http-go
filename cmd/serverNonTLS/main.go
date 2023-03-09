package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	fmt.Println("Server listening on port 9080...")
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
