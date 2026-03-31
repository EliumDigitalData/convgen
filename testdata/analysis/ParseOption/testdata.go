//go:build convgen

package testdata

import "github.com/EliumDigitalData/convgen"

var renameToLower = convgen.RenameToLower(true, true) // want `cannot assign RenameToLower to variable`

func asis[T any](T) T { return *new(T) }

var m = (convgen.Module(
	renameToLower, // want `option must be inlined, not assigned to variable`

	asis(convgen.RenameToLower(true, true)), // want `option must be convgen directive`

	convgen.RenameToLower(true, true),   // ok
	(convgen.RenameToLower(true, true)), // ok
))
