//go:build convgen

package main

import "github.com/EliumDigitalData/convgen"

var Good = convgen.Struct[struct{}, *struct{}](nil)
