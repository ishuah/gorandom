package main

import (
	"flag"
	"fmt"
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

	response, err := Get(length, dataType, size)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Data)
	return
}
