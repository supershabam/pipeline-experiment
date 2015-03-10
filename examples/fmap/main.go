package main

import (
	"fmt"
	"strconv"
)

func transform(in int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < in; i++ {
			// time.Sleep(500 * time.Millisecond) // simulate delay
			out <- strconv.Itoa(in)
		}
	}()
	return out
}

//go:generate pipeliner -name fmap -fn fmapInt -in int -out string -package main -file fmapInt.go
// pipeliner flatMap(int) <-chan string concurrently
func main() {
	input := make(chan int)
	go func() {
		defer close(input)
		for i := 0; i < 5; i++ {
			input <- i
		}
	}()
	output := fmapInt(10, transform, input)
	for o := range output {
		fmt.Printf("out: %s\n", o)
	}
}
