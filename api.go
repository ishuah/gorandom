package main

import(
	"io"
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
)

var URL = "https://qrng.anu.edu.au/API/jsonI.php"

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
	Data []int `json: "data"`
	Success bool `json: "success"`
}

func processHex(body io.ReadCloser) (output string, err error) {
	jsonResponse := HexResponse{}
	err = json.NewDecoder(body).Decode(&jsonResponse)
	if err != nil {
		return
	}

	for i := range jsonResponse.Data {
		output = output + " " + jsonResponse.Data[i]
	}

	return
}

func processInt(body io.ReadCloser) (output string, err error) {
	jsonResponse := IntResponse{}
	err = json.NewDecoder(body).Decode(&jsonResponse)
	if err != nil {
		return
	}

	for i := range jsonResponse.Data {
		output = output + " " + strconv.Itoa(jsonResponse.Data[i])
	}
	return
}

func Get(length int, dataType string, size int) (output string, err error) {
	URLWithParams := URL + fmt.Sprintf("?length=%v&type=%v&size=%v", length, dataType, size)
	resp, err := http.Get(URLWithParams)

	if err != nil {
		return
	}

	if (dataType == "hex16") {
		output, err = processHex(resp.Body)
	} else {
		output, err = processInt(resp.Body)
	}
	return
}