//go:build convgen

package main

import (
	"fmt"

	"github.com/sublee/convgen"
)

type (
	AuditData struct {
		CreatedAt int
		UpdatedAt int
	}

	A struct {
		Name      string
		CreatedAt int
		UpdatedAt int
	}

	B struct {
		Name string
		AuditData
	}
)

var mod = convgen.Module(convgen.MatchEmbedded(1))

var A2B = convgen.Struct[A, B](mod)
var B2A = convgen.Struct[B, A](mod)

func main() {
	a := A{Name: "x", CreatedAt: 1, UpdatedAt: 2}
	b := A2B(a)
	fmt.Println(b.Name, b.AuditData.CreatedAt, b.AuditData.UpdatedAt)

	b2 := B{Name: "y", AuditData: AuditData{CreatedAt: 3, UpdatedAt: 4}}
	a2 := B2A(b2)
	fmt.Println(a2.Name, a2.CreatedAt, a2.UpdatedAt)
}
