//go:build convgen

package main

import (
	"github.com/EliumDigitalData/convgen"
)

type (
	InnerA struct{ ID int }
	InnerB struct{ ID int }

	// Src has ID reachable via two paths at depth 1 — ambiguous.
	Src struct {
		InnerA
		InnerB
	}

	Dst struct{ ID int }
)

var conv = convgen.Struct[Src, Dst](nil, convgen.MatchEmbedded(1))

func main() {}
