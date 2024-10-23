// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate stringer -type=Size -linecomment -output size_generated.go
package main

// Size options. MD is default in DaisyUI so it goes first.
const (
	MD Size = iota // md
	XS             // xs
	SM             // sm
	LG             // lg
)

type Size int

func AddSize(object string, size Size) string {
	return object + "-" + size.String()
}
