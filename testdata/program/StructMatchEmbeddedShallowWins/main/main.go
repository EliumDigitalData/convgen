//go:build convgen

package main

import (
	"fmt"

	"github.com/sublee/convgen"
)

// X has Value at both depth 0 and depth 1 (inside Inner).
// The depth-0 Value should win on the X side.
type (
	InnerX struct{ Value int }
	X      struct {
		Value int // depth 0 — shadows InnerX.Value
		InnerX
	}
	Y struct{ Value int }
)

var mod = convgen.Module(convgen.MatchEmbedded(1))

var X2Y = convgen.Struct[X, Y](mod)
var Y2X = convgen.Struct[Y, X](mod)

func main() {
	x := X{Value: 7, InnerX: InnerX{Value: 99}}
	y := X2Y(x)
	fmt.Println(y.Value) // should be 7 (depth-0 wins)

	y2 := Y{Value: 5}
	x2 := Y2X(y2)
	fmt.Println(x2.Value) // should be 5
}
