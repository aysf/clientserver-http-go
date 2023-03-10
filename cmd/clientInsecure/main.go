package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// resp, err := client.Get("http://localhost:9080/")

	dt := make(map[string]interface{})
	dt["Name"] = "Ruby Ad"

	// payload := new(bytes.Buffer)
	payload := &bytes.Buffer{}

	json.NewEncoder(payload).Encode(dt)

	resp, err := client.Post("https://localhost:9080/data", "application/json", payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// cara 1

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("ERROR: read response body:", err)
	// 	return
	// }
	// fmt.Println("result body:", string(body))

	// cara 2

	r := new(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		log.Fatal("ERROR: failed decode response:", err)
		return
	}
	fmt.Println("result body:", r)

	fmt.Println("result code:", resp.StatusCode)
}
