package main

import "fmt"

func main() {
	input := make(chan int)
	go func() {
		defer close(input)
		input <- 1
		input <- 2
		input <- 3
	}()
	output := scan(input)
	for o := range output {
		fmt.Printf("%d\n", o)
	}
}

// scan looks at the last two elements in a channel and produces a result from
// them
func scan(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		last := <-in
		for curr := range in {
			out <- last + curr
			last = curr
		}
	}()
	return out
}
