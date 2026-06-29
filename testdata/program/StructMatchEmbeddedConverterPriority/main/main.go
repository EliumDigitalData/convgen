//go:build convgen

package main

import (
	"fmt"
	"time"

	"github.com/EliumDigitalData/convgen"
)

// AuditData and AuditDTO have different field names.
// Without converter priority, MatchEmbedded would try to promote CreatedAt/UpdatedAt
// from AuditData and Created/Updated from AuditDTO — no name matches, so it would fail.
// With converter priority, AuditData/AuditDTO are kept as leaves and AuditConv is used.
type AuditData struct {
	CreatedAt int
	UpdatedAt int
}

type AuditDTO struct {
	Created int
	Updated int
}

type Time1 time.Time
type Time2 time.Time

type B struct {
	Name string
	AuditData
	Time Time1
}

type C struct {
	Name string
	AuditDTO
	Time Time2
}

var mod = convgen.Module(
	convgen.MatchEmbedded(1),
	convgen.ImportFunc(func(t1 Time1) Time2 { return Time2(t1) }),
)

// Explicit converter for the embedded struct pair.
// Its presence causes MatchEmbedded to treat AuditData (as input) and AuditDTO
// (as output) as opaque leaves rather than promoting their sub-fields.
var AuditConv = convgen.Struct[AuditData, AuditDTO](mod,
	convgen.Match(AuditData{}.CreatedAt, AuditDTO{}.Created),
	convgen.Match(AuditData{}.UpdatedAt, AuditDTO{}.Updated),
)

// B2C: B.AuditData is a leaf (AuditConv converts FROM AuditData),
//
//	C.AuditDTO is a leaf (AuditConv converts TO AuditDTO).
//
// An explicit Match links the two leaf fields; AuditConv handles the conversion.
var B2C = convgen.Struct[B, C](mod,
	convgen.Match(B{}.AuditData, C{}.AuditDTO),
)

func main() {
	b := B{Name: "hello", AuditData: AuditData{CreatedAt: 10, UpdatedAt: 20}, Time: Time1(time.Unix(100, 0))}
	c := B2C(b)
	fmt.Println(c.Name, c.AuditDTO.Created, c.AuditDTO.Updated, time.Time(c.Time).Unix())
}
