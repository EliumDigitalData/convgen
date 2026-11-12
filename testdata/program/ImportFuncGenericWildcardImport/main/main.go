//go:build convgen

package main

import (
	"fmt"

	"github.com/EliumDigitalData/convgen"

	"example.com/ImportFuncGenericWildcardImport/box"
	"example.com/ImportFuncGenericWildcardImport/money"
)

// Foo and Bar only ever mention the money package through the field type
// box.Box[money.Money]/box.Wrap[money.Money]. The generated converter calls
// box.Convert[money.Money](in.Value), where money.Money only appears as the
// solved type argument of the wildcard call — nowhere else in the generated
// file — so this specifically exercises that the generated code registers an
// import for money even though it's never otherwise referenced.
type Foo struct {
	Value box.Box[money.Money]
}

type Bar struct {
	Value box.Wrap[money.Money]
}

var mod = convgen.Module(convgen.ImportFunc(box.Convert[convgen.G]))

var Conv = convgen.Struct[Foo, Bar](mod)

func main() {
	out := Conv(Foo{Value: box.Box[money.Money]{V: money.Money{Cents: 150}}})
	fmt.Println(out.Value.V.Cents)
}
