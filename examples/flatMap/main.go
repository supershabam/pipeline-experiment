package main

import "fmt"

// example of a flatMap type pattern where a channel of channels is turned into
// a single output channel
func main() {
	for i := range flatMap(gen()) {
		fmt.Printf("%d\n", i)
	}
}

func gen() <-chan (<-chan int) {
	out := make(chan (<-chan int))
	go func() {
		defer close(out)
		out <- intRange()
		out <- intRange()
		out <- intRange()
	}()
	return out
}

func intRange() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		out <- 1
		out <- 2
		out <- 3
	}()
	return out
}

func flatMap(in <-chan (<-chan int)) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for c := range in {
			for i := range c {
				out <- i
			}
		}
	}()
	return out
}
