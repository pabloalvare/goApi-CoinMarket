package main

import (
	"io/ioutil"
	"log"
)

func GetApiKey() string {
	file, err := ioutil.ReadFile("apiKey.txt")
	if err != nil {
		log.Fatal(err)
	}
	key := string(file)
	return key

}
