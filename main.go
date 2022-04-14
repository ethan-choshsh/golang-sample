package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	statusCode, header := GetPolarisshare()
	fmt.Printf("상태코드: %d\n헤더: %s", statusCode, header)
}

func GetPolarisshare() (int, string) {
	// HTTP GET
	res, err := http.Get("https://polarishare.com/new")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(112312312312)

	// Print header
	header, err := json.MarshalIndent(res.Header, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return res.StatusCode, string(header)
}
