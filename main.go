package main

import (
	"fmt"
	"log"
	"flag"
)

func main() {
	var length int
	var dataType string
	var size int

	flag.IntVar(&length, "length", 1, "Array length")
	flag.StringVar(&dataType, "type", "hex16", "Data type")
	flag.IntVar(&size, "size", 5, "Block size")
	flag.Parse()

	
	output, err := Get(length, dataType, size)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
	return
}