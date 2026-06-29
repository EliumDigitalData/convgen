//go:build convgen

package main

import (
	"fmt"

	"github.com/EliumDigitalData/convgen"
)

// Timestamp is a struct type we want to keep opaque — its internal fields
// should NOT be promoted into the matchable set.
type Timestamp struct {
	Seconds int
	Nanos   int
}

type (
	// Event has a Timestamp field and a nested metadata struct.
	Event struct {
		Metadata struct {
			ID int
		}
		OccurredAt Timestamp
	}

	// Flat represents the promoted view of Event.
	Flat struct {
		ID         int
		OccurredAt Timestamp
	}
)

var mod = convgen.Module(
	convgen.MatchEmbedded(1),
	convgen.MatchEmbeddedExclude[Timestamp](),
)

// Event.Metadata.ID -> Flat.ID   (promoted from depth 1)
// Event.OccurredAt  -> Flat.OccurredAt (opaque leaf at depth 0; Seconds/Nanos NOT promoted)
var Event2Flat = convgen.Struct[Event, Flat](mod)
var Flat2Event = convgen.Struct[Flat, Event](mod)

func main() {
	e := Event{
		Metadata:   struct{ ID int }{ID: 42},
		OccurredAt: Timestamp{Seconds: 1000, Nanos: 500},
	}
	f := Event2Flat(e)
	fmt.Println(f.ID, f.OccurredAt.Seconds, f.OccurredAt.Nanos)

	f2 := Flat{ID: 7, OccurredAt: Timestamp{Seconds: 2000, Nanos: 0}}
	e2 := Flat2Event(f2)
	fmt.Println(e2.Metadata.ID, e2.OccurredAt.Seconds, e2.OccurredAt.Nanos)
}
