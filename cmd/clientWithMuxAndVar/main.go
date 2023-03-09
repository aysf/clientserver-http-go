package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	URL1 = "http://localhost:9080/one"
	URL2 = "http://localhost:9080/viewquery?param1=abc&param2=234"
	URL3 = "http://localhost:9080/user/123"
)

func main() {

	// br := "name=anan&age=22&date=22121991"

	// req, _ := http.NewRequest("GET", URL, nil)
	req, err := CreateRequest("POST", URL3, nil)
	if err != nil {
		log.Fatal("ERORR: create request:", err)
	}

	req.URL.RawQuery = "numbers=1,2,3&param1=random233"

	client := &http.Client{}

	res, _ := client.Do(req)

	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
	fmt.Printf("resut: %#v, type: %T", res.Body, res.Body)

	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(body)

}

func CreateRequest(method string, url string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return req, err
}
