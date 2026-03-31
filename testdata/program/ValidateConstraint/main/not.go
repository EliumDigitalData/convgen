//go:build !convgen

package main

import "github.com/EliumDigitalData/convgen"

var Not = convgen.Struct[struct{}, *struct{}](nil)
