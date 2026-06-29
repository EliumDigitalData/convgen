//go:build convgen

package main

import (
	"fmt"

	"github.com/sublee/convgen"
)

type (
	Inner struct{ Value int }
	Mid   struct{ Inner }
	Deep  struct{ Mid }

	Flat struct{ Value int }
)

var mod = convgen.Module(convgen.MatchEmbedded(2))

// Deep.Mid.Inner.Value (depth 2) -> Flat.Value (depth 0)
var Deep2Flat = convgen.Struct[Deep, Flat](mod)
var Flat2Deep = convgen.Struct[Flat, Deep](mod)

func main() {
	d := Deep{Mid: Mid{Inner: Inner{Value: 42}}}
	f := Deep2Flat(d)
	fmt.Println(f.Value)

	f2 := Flat{Value: 99}
	d2 := Flat2Deep(f2)
	fmt.Println(d2.Mid.Inner.Value)
}
