package main

import (
	"io/ioutil"
	"log"
)

const src = `
package main

import "fmt"

func main() {
	fmt.Println("go pipeline something else ya jerk")
}
`

func main() {
	err := ioutil.WriteFile("./pipelined.go", []byte(src), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
