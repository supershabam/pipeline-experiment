package main

import (
	"fmt"
	"strconv"
)

//go:generate pipeliner -name batch -fn batchString -in string -file batchString.go -package main
func main() {
	input := make(chan string)
	go func() {
		defer close(input)
		for i := 0; i < 37; i++ {
			input <- strconv.Itoa(i)
		}
	}()
	batch := batchString(10, input)
	for output := range batch {
		fmt.Printf("batch size: %d\n", len(output))
	}
}
