package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// URL points to the Quantum Random Number Generator API
var URL = "https://qrng.anu.edu.au/API/jsonI.php"

// Response describes the response from the qrng API
type Response struct {
	DataType string      `json:"type"`
	Length   int         `json:"length"`
	Size     int         `json:"size"`
	Data     interface{} `json:"data"`
	Success  bool        `json:"success"`
}

// Get makes a formatted GET request with the parameters supplied
func Get(length int, dataType string, size int) (jsonResponse Response, err error) {
	URLWithParams := URL + fmt.Sprintf("?length=%v&type=%v&size=%v", length, dataType, size)
	resp, err := http.Get(URLWithParams)

	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)

	return
}
