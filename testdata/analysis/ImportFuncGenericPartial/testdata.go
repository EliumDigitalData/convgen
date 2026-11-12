//go:build convgen

package testdata

import "github.com/EliumDigitalData/convgen"

func Pair[K comparable, V any](k K) map[K]V { return nil }

var _ = convgen.Module(
	convgen.ImportFunc(Pair[convgen.G, int]), // want `convgen.G must be used for every type parameter`
)
