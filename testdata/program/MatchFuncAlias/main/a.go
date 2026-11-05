//go:build convgen

package main

import (
	extA "example.com/MatchFuncAlias/ext"

	"github.com/EliumDigitalData/convgen"
)

type X struct{ C extA.Config }
type Y struct{ C int }

var ConvA = convgen.Struct[X, Y](nil,
	convgen.MatchFunc(X{}.C, Y{}.C, func(c extA.Config) int { return c.N }),
)
