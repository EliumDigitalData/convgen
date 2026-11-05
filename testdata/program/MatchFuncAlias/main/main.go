package main

import (
	"fmt"

	"example.com/MatchFuncAlias/ext"
)

func main() {
	a := ConvA(X{C: ext.Config{N: 3}})
	b := ConvB(P{C: ext.Config{N: 5}})

	// Output: 3 10
	fmt.Println(a.C, b.C)
}
