# grand
A tool for interfacing with the ANU Quantum Random Number Generator

## installation
This project does not have release builds yet.
To run it locally, clone this repo and run `go build` in the project directory.

## usage
Grand takes three parameters:
 - `--type` specifies the type of random number returned. Options are `uint8`, `uint16` and `hex16`. Defaults to `hex16`
 - `--length` specifies the number of random numbers returned. Max value: 1024. Defaults to 1
 - `--size` specifies the length of hex16 blocks. Defaults to 12

 Example:
 ```bash
 $ grand --size 32
 7ebf0eaba4c2f19708fe0d7dbb54d66af3dbf3fce9e9e66621a1890d0a1d539d
 ```