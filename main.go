package main

import (
	"fmt"
	"os"
	"perceptual_hash/hashes"
)

func main() {
	path := os.Args[1]
	hash := hashes.CalcDHash(path)
	fmt.Printf("hash is 0b%b\n", hash)
	fmt.Printf("hash is 0x%X\n", hash)
}
