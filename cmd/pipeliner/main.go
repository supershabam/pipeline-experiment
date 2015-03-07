package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/supershabam/pipeline"
)

// TODO make parameters to pipeliner much cleaner e.g. `pipeliner batch(string)`
var (
	name = flag.String("name", "", "name of pipeline to generate")
	in   = flag.String("in", "", "in type")
	out  = flag.String("out", "", "out type")
	pkg  = flag.String("package", "", "package of file")
	fn   = flag.String("fn", "", "name of function")
	file = flag.String("file", "", "output file")
)

func main() {
	flag.Parse()
	f, err := os.Create(*file)
	if err != nil {
		log.Fatal(err)
	}
	switch *name {
	case "batch":
		err = pipeline.RenderBatch(f, pipeline.BatchConfig{
			Package:   *pkg,
			FuncName:  *fn,
			Type:      *in,
			Timestamp: time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"),
			Version:   version,
		})
	default:
		log.Fatalf("unknown pipeliner type: %s", *name)
	}
}
