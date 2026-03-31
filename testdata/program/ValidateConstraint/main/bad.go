package main

// Convgen is imported but this file has no "//go:build convgen" constraint.
import "github.com/EliumDigitalData/convgen"

var Bad = convgen.Struct[struct{}, *struct{}](nil)
