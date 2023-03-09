package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 9080...")
	err := http.ListenAndServeTLS(":9080", "../../server.crt", "../../server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
