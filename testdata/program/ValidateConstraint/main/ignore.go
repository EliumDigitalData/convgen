//go:build ignore

package main

import "github.com/EliumDigitalData/convgen"

var Ignore = convgen.Struct[struct{}, *struct{}](nil)
