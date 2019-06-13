package main

import (
	"fmt"
	"flag"
	"net/http"
	"encoding/json"
)

var URL = "https://qrng.anu.edu.au/API/jsonI.php?length=5&type=uint8&size=8"

type HexResponse struct {
	DataType string `json: "type"`
	Length int `json: "length"`
	Size int `json: "size"`
	Data []string `json: "data"`
	Success bool `json: "success"`
}

type IntResponse struct {
	DataType string `json: "type"`
	Length int `json: "length"`
	Size int `json: "size"`
	Data []uint `json: "data"`
	Success bool `json: "success"`
}

func main() {
	var length int
	var dataType string
	var size int

	flag.IntVar(&length, "length", 1, "Array length")
	flag.StringVar(&dataType, "type", "hex16", "Data type")
	flag.IntVar(&size, "size", 5, "Block size")
	flag.Parse()

	resp, err := http.Get(URL + fmt.Sprintf("?length=%v&type=%v&size=%v", length, dataType, size))

	if err != nil {
		fmt.Printf("Error making GET request: %v \n", err)
		return
	}

	jsonResponse := HexResponse{}
	

	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		fmt.Printf("Error decoding response: %v \n", err)
		return
	}

	fmt.Printf("%v\n", jsonResponse.Data)
	return
}