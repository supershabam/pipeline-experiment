package pipeline

import (
	"fmt"
	"go/ast"
	"go/parser"
	"io"
	"strings"
)

type TransformFunc struct {
	In  string
	Out string
}

// Parse expects to parse a golang function definition where there is a single
// input type and a single output type
func (tf *TransformFunc) Parse(x string) error {
	t, err := parser.ParseExpr(x)
	if err != nil {
		return err
	}
	var ft *ast.FuncType
	var ok bool
	if ft, ok = t.(*ast.FuncType); !ok {
		return fmt.Errorf("expected func definition")
	}
	if ft.Params.NumFields() != 1 {
		return fmt.Errorf("expected one field")
	}
	if ft.Results.NumFields() != 1 {
		return fmt.Errorf("expected one return")
	}
	tf.In = x[ft.Params.List[0].Type.Pos()-1 : ft.Params.List[0].Type.End()-1]
	tf.Out = x[ft.Results.List[0].Type.Pos()-1 : ft.Results.List[0].Type.End()-1]
	return nil
}

type Parser struct {
	Version string
}

type MapPipeliner struct {
	TF TransformFunc
}

func (mp MapPipeliner) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

func (p Parser) Parse(x string) (Pipeliner, error) {
	name := strings.SplitN(x, "(", 2)[0]
	switch name {
	case "map":
		return p.ParseMap(x)
	}
	return nil, fmt.Errorf("not implemented")
}

// ParseMap expects "map(func(T) Q) [concurrently]"
func (p Parser) ParseMap(x string) (MapPipeliner, error) {
	start := len("map(")
	lparens := 0
	end := start
	for ; end < len(x)-1; end++ {
		if x[end+1] == '(' {
			lparens++
		} else if x[end+1] == ')' && lparens == 0 {
			break
		}
	}
	transform := x[start:end]
	var tf TransformFunc
	err := tf.Parse(transform)
	if err != nil {
		return MapPipeliner{}, err
	}
	return MapPipeliner{TF: tf}, nil
}
