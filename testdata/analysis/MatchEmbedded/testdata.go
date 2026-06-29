//go:build convgen

package testdata

import (
	"github.com/EliumDigitalData/convgen"
)

type (
	embX  struct{ A int }
	embY  struct{ A int }
	embSt struct{ Z int }
)

// valid depth
var Emb1 = convgen.Struct[embX, embY](nil,
	convgen.MatchEmbedded(1), // ok
)

// depth 0 is valid (explicitly disables)
var Emb2 = convgen.Struct[embX, embY](nil,
	convgen.MatchEmbedded(0), // ok
)

// negative depth is not allowed
var Emb3 = convgen.Struct[embX, embY](nil,
	convgen.MatchEmbedded(-1), // want `MatchEmbedded depth must be non-negative`
)

// exclude with a type argument is ok
var Emb4 = convgen.Struct[embX, embY](nil,
	convgen.MatchEmbedded(1),
	convgen.MatchEmbeddedExclude[embSt](), // ok
)
