//go:build convgen

package main

import (
	"fmt"

	"github.com/sublee/convgen"
)

// AuditData is a named embedded struct.
type AuditData struct {
	CreatedAt int
	UpdatedAt int
}

// B embeds AuditData anonymously.
type B struct {
	Name string
	AuditData
}

// C has a named nested struct field and a direct field sharing the same names
// as AuditData fields.
type C struct {
	Name  string
	Audit struct {
		CreatedAt int
	}
	UpdatedAt int
}

var mod = convgen.Module(convgen.MatchEmbedded(1))

// B.AuditData.CreatedAt -> C.Audit.CreatedAt (both at depth 1)
// B.AuditData.UpdatedAt -> C.UpdatedAt       (B depth 1, C depth 0 — shallowest wins)
var B2C = convgen.Struct[B, C](mod)
var C2B = convgen.Struct[C, B](mod)

func main() {
	b := B{Name: "hello", AuditData: AuditData{CreatedAt: 10, UpdatedAt: 20}}
	c := B2C(b)
	fmt.Println(c.Name, c.Audit.CreatedAt, c.UpdatedAt)

	c2 := C{Name: "world", Audit: struct{ CreatedAt int }{CreatedAt: 30}, UpdatedAt: 40}
	b2 := C2B(c2)
	fmt.Println(b2.Name, b2.AuditData.CreatedAt, b2.AuditData.UpdatedAt)
}
