// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Size -linecomment -output size_generated.go
package components

// Size options. MD is default in DaisyUI so it goes first.
const (
	MD Size = iota // md
	XS             // xs
	SM             // sm
	LG             // lg
	XL             // xl
)

type Size int

// SizeObject returns the given object class with an appropriate size
// modifier. For e.g., if Size is SM and object is menu, the return
// value will be "menu-sm".
func (s Size) SizeObject(object string) string {
	return object + "-" + s.String()
}
