//go:build convgen

package main

import (
	"example.com/MatchFuncAlias/ext"

	"github.com/EliumDigitalData/convgen"
)

type P struct{ C ext.Config }
type Q struct{ C int }

var ConvB = convgen.Struct[P, Q](nil,
	convgen.MatchFunc(P{}.C, Q{}.C, func(c ext.Config) int { return c.N * 2 }),
)
