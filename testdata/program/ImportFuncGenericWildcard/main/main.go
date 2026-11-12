//go:build convgen

package main

import (
	"fmt"

	"github.com/EliumDigitalData/convgen"
)

type Box[T any] struct {
	V T
}

type Wrap[T any] struct {
	V T
}

func box2wrap[T any](b Box[T]) Wrap[T] {
	return Wrap[T]{V: b.V}
}

type Crate[T any] struct {
	V T
}

type Parcel[T any] struct {
	V T
}

func crate2parcel[T any](c Crate[T]) (Parcel[T], error) {
	return Parcel[T]{V: c.V}, nil
}

var mod = convgen.Module(
	convgen.ImportFunc(box2wrap[convgen.G]),
	convgen.ImportFuncErr(crate2parcel[convgen.G]),
)

type FooStr struct {
	Value Box[string]
}

type BarStr struct {
	Value Wrap[string]
}

type FooInt struct {
	Value Box[int]
}

type BarInt struct {
	Value Wrap[int]
}

type BazStr struct {
	Value Crate[string]
}

type QuxStr struct {
	Value Parcel[string]
}

// ConvStr and ConvInt both resolve through the SAME wildcard registration
// (box2wrap[convgen.G]), instantiated at different concrete T.
var ConvStr = convgen.Struct[FooStr, BarStr](mod)
var ConvInt = convgen.Struct[FooInt, BarInt](mod)

var ConvErr = convgen.StructErr[BazStr, QuxStr](mod)

func main() {
	str := ConvStr(FooStr{Value: Box[string]{V: "42"}})
	fmt.Println(str.Value.V)

	i := ConvInt(FooInt{Value: Box[int]{V: 7}})
	fmt.Println(i.Value.V)

	q, _ := ConvErr(BazStr{Value: Crate[string]{V: "hi"}})
	fmt.Println(q.Value.V)
}
