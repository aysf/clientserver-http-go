package main

import (
	"fmt"
	"net/http"
)

func main() {

	// URL := "http://localhost:9080/one"
	URL := "http://localhost:9080/cekquery?param1=abc&param2=234"

	req, _ := http.NewRequest("GET", URL, nil)

	client := &http.Client{}

	res, _ := client.Do(req)

	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
	fmt.Printf("resut: %#v", res.Body)

	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(body)

}
