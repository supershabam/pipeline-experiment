package main

import (
	"log"
	"os"
	"time"

	"github.com/supershabam/pipeline"
)

func main() {
	file, err := os.Create("./pipeline_maxSizeBatchString.go")
	if err != nil {
		log.Fatal(err)
	}
	err = pipeline.RenderMaxSizeBatch(file, pipeline.MaxSizeBatchConfig{
		Package:   "main",
		FuncName:  "batchString",
		Type:      "string",
		Timestamp: time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"),
	})
	if err != nil {
		log.Fatal(err)
	}
}
