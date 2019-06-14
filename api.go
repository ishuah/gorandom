package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

var URL = "https://qrng.anu.edu.au/API/jsonI.php"

type Response struct {
	DataType string `json: "type"`
	Length int `json: "length"`
	Size int `json: "size"`
	Data interface{} `json: "data"`
	Success bool `json: "success"`
}

func Get(length int, dataType string, size int) (jsonResponse Response, err error) {
	URLWithParams := URL + fmt.Sprintf("?length=%v&type=%v&size=%v", length, dataType, size)
	resp, err := http.Get(URLWithParams)

	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	
	return
}