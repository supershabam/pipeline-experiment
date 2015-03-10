package pipeline

import "io"

type Pipeliner interface {
	io.WriterTo
}
