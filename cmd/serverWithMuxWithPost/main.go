package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

type Person struct {
	Name      string `json:"name"`
	Age       uint   `json:"age"`
	IsMarried bool   `json:"is_married"`
	Password  string `json:"-"`
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/person", http.HandlerFunc(person))

	fmt.Println("Server listening on port 9080...")
	l, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatal("error net listen:", err)
	}

	http.Serve(l, mux)

}

func person(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		log.Fatal("method not allowed")
	}

	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "error read request", http.StatusInternalServerError)
	// }

	p := &Person{}

	// by := bytes.NewBuffer(b)

	j := json.NewDecoder(r.Body)
	err := j.Decode(p)
	if err != nil {
		http.Error(w, "error decode request", http.StatusInternalServerError)
	}

	log.Printf("result: %+v", p)

	data := map[string]interface{}{
		"success": true,
		"data":    p,
	}

	je := json.NewEncoder(w)
	err = je.Encode(data)
	if err != nil {
		http.Error(w, "error encode response", http.StatusInternalServerError)
	}

}
