package main

import (
	"fmt"
	"strconv"
	"time"
)

func transform(in int) string {
	time.Sleep(500 * time.Millisecond) // simulate delay
	return strconv.Itoa(in)
}

//go:generate pipeliner -name cmap -fn cmapInt -in int -out string -package main -file cmapInt.go
func main() {
	input := make(chan int)
	go func() {
		defer close(input)
		for i := 0; i < 25; i++ {
			input <- i
		}
	}()
	output := cmapInt(10, transform, input)
	for o := range output {
		fmt.Printf("out: %s\n", o)
	}
}
