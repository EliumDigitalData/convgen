//go:build convgen

package main

import (
	"fmt"

	"github.com/EliumDigitalData/convgen"
)

// Nested is embedded behind a pointer in A. An ImportFunc is registered for
// *Nested, not Nested. MatchEmbedded must still recognize this as an existing
// converter and keep A.Field as a leaf instead of promoting Nested.OtherField,
// even though the traversal compares against the fully dereferenced type.
type Nested struct {
	OtherField string
}

type A struct {
	Field *Nested
}

type B struct {
	Field string
}

var mod = convgen.Module(
	convgen.MatchEmbedded(1),
	convgen.ImportFunc(func(n *Nested) string { return n.OtherField }),
)

var A2B = convgen.Struct[A, B](mod)

func main() {
	b := A2B(A{Field: &Nested{OtherField: "hello"}})
	fmt.Println(b.Field)
}
