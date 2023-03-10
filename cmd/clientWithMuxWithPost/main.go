package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	URL1 = "http://localhost:9080/person"
)

type Person struct {
	Name      string `json:"name"`
	Age       uint   `json:"age"`
	IsMarried bool   `json:"is_married"`
	Password  string `json:"password"`
}

type Response struct {
	Success bool `json:"success"`
	Data    Person
}

var (
	// data json
	PersonA = &Person{"Jack", 33, true, "supersecret"}
)

func main() {

	p, err := json.Marshal(PersonA)
	if err != nil {
		log.Fatalln("ERRO: marshal person:", err)
	}

	req, err := CreateRequest("POST", URL1, strings.NewReader(string(p)))
	if err != nil {
		log.Fatal("ERORR: create request:", err)
	}
	// req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, _ := client.Do(req)

	fmt.Println(res.StatusCode)

	r := &Response{}
	json.NewDecoder(res.Body).Decode(r)

	fmt.Printf("response json: %+v\n", r)

}

func CreateRequest(method string, url string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return req, err
}
