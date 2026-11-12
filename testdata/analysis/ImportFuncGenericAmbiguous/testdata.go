//go:build convgen

package testdata

import "github.com/EliumDigitalData/convgen"

type Box[T any] struct {
	V T
}

func boxToBox1[T any](b Box[T]) Box[T] { return b }
func boxToBox2[T any](b Box[T]) Box[T] { return b }

var mod = convgen.Module(
	convgen.ImportFunc(boxToBox1[convgen.G]),
	convgen.ImportFunc(boxToBox2[convgen.G]),
)

type Foo struct {
	Value Box[string]
}

type Bar struct {
	Value Box[string]
}

var Conv = convgen.Struct[Foo, Bar](mod) // want `ambiguous generic converters`
