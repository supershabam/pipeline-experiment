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
	//go:generate pipeline scan([]int{}, lastTwo) as windowScan into window_scan.go
	output := windowScan(input)
	for o := range output {
		fmt.Printf("%d\n", o)
	}
}

// defined in the package to tell windowScan what types to deal with
func lastTwo(acc []int, x int) []int {
	acc = append(acc, x)
	if len(acc) > 2 {
		acc = acc[1:]
	}
	return acc
}

// automatically generated
func windowScan(in <-chan int) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		acc := []int{}
		for x := range in {
			acc = lastTwo(acc, x)
			out <- acc
		}
	}()
	return out
}
