package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	input := make(chan string)
	go func() {
		defer close(input)
		for i := 0; i < 20; i++ {
			input <- fmt.Sprintf("%d", i)
		}
	}()

	for output := range concMap(fOfS, input) {
		fmt.Printf("output: %s\n", output)
	}
}

func fOfS(in string) string {
	time.Sleep(500 * time.Millisecond)
	return fmt.Sprintf("fOfS(%s)", in)
}

func concMap(fn func(string) string, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		wg := sync.WaitGroup{}
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for item := range in {
					out <- fOfS(item)
				}
			}()
		}
		wg.Wait()
	}()
	return out
}
