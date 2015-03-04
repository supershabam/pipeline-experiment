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
		input <- "too long for mah function"
	}()

	// for output := range concMap(20, fOfS, input) {
	// fmt.Printf("output: %s\n", output)
	// }

	output, errc := concErrMap(20, fOfSErr, input)
	for out := range output {
		fmt.Printf("output: %s\n", out)
	}
	err := <-errc
	if err != nil {
		fmt.Print(err)
	}
}

func fOfS(in string) string {
	time.Sleep(500 * time.Millisecond)
	return fmt.Sprintf("fOfS(%s)", in)
}

func fOfSErr(in string) (string, error) {
	time.Sleep(100 * time.Millisecond)
	if len(in) > 10 {
		return "", fmt.Errorf("string too long!")
	}
	return fmt.Sprintf("fOfSErr(%s)", in), nil
}

//go:generate pipeliner map(func(string) string) concurrently as concMap into conc_map.go
func concMap(concurrency int, fn func(string) string, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		wg := sync.WaitGroup{}
		wg.Add(concurrency)
		for i := 0; i < concurrency; i++ {
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

//go:generate pipeliner map(func(string) (string, error) concurrently as concErrMap into conc_err_map.go
func concErrMap(concurrency int, fn func(string) (string, error), in <-chan string) (<-chan string, <-chan error) {
	out := make(chan string)
	errc := make(chan error, 1)
	done := make(chan struct{})
	once := sync.Once{}
	go func() {
		defer close(out)
		var outerErr error
		defer func() {
			errc <- outerErr
		}()
		wg := sync.WaitGroup{}
		wg.Add(concurrency)
		for i := 0; i < concurrency; i++ {
			go func() {
				defer wg.Done()
				for {
					select {
					case <-done:
						return
					case item, ok := <-in:
						if !ok {
							return // end of channel
						}
						t, err := fn(item)
						if err != nil {
							once.Do(func() {
								outerErr = err
								close(done)
							})
							return
						}
						out <- t
					}
				}
			}()
		}
		wg.Wait()
	}()
	return out, errc
}
