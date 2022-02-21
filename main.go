package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := *&http.Client{}

	req, err := GetRequest()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(respBody))

}
