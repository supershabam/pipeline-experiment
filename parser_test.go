package pipeline

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	p := Parser{Version: "2"}
	pl, err := p.Parse("map(func(<-chan (<-chan string)) (string, error))")
	if err != nil {
		t.Error(err)
		return
	}
	if _, ok := pl.(MapPipeliner); !ok {
		t.Errorf("expected pl to be a map pipeliner")
	}
	fmt.Printf("%+v", pl)
}
