package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int32
}

type Q struct {
	X, Y *int32
}

func main() {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	err := enc.Encode(P{3, 4, 5})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	fmt.Println(network)

	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(*q.X, *q.Y)
}
