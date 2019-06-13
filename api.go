package main

import(
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
)

var URL = "https://qrng.anu.edu.au/API/jsonI.php"

type Response interface {
	PrintData() (string, error)
}

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

func (r *HexResponse) PrintData() (output string, err error) {
	for i := range r.Data {
		output = output + " " + r.Data[i]
	}
	return
}

func (r *IntResponse) PrintData() (output string, err error) {
	for i := range r.Data {
		output = output + " " + strconv.Itoa(r.Data[i])
	}
	return
}

func Get(length int, dataType string, size int) (jsonResponse Response, err error) {
	URLWithParams := URL + fmt.Sprintf("?length=%v&type=%v&size=%v", length, dataType, size)
	resp, err := http.Get(URLWithParams)

	if err != nil {
		return
	}
	
	if (dataType == "hex16") {
		jsonResponse = &HexResponse{}
	} else {
		jsonResponse = &IntResponse{}
	}

	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	
	return
}