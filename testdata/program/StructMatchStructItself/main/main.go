//go:build convgen

package main

import (
	"github.com/EliumDigitalData/convgen"
)

type (
	X struct{}
	Y struct{}
)

var XtoY = convgen.StructErr[X, Y](nil,
	convgen.Match(X{}, Y{}),
)

func main() {
	panic("convgen will fail")
}
