//go:build convgen && !negvnoc

package main

import "github.com/EliumDigitalData/convgen"

var And = convgen.Struct[struct{}, *struct{}](nil)
