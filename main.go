package main

import (
	"flag"
	"fmt"
	"github.com/ishuah/gorandom/api"
	"log"
)

func main() {
	var length int
	var dataType string
	var size int

	flag.IntVar(&length, "length", 1, "Array length")
	flag.StringVar(&dataType, "type", "hex16", "Data type")
	flag.IntVar(&size, "size", 5, "Block size")
	flag.Parse()

	response, err := api.Get(length, dataType, size)

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range response.Data {
		fmt.Println(value)
	}

	return
}
