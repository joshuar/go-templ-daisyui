// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Alignment -linecomment -output alignment_generated.go
package components

const (
	AlignEnd    Alignment = iota // end
	AlignTop                     // top
	AlignBottom                  // bottom
	AlignMiddle                  // middle
	AlignLeft                    // left
	AlignRight                   // right
)

type Alignment int

// AlignObject returns the given object class with an appropriate alignment
// modifier. For e.g., if Alignment is AlignTop and object is menu, the return
// value will be "menu-top".
func (a Alignment) AlignObject(object string) string {
	return object + "-" + a.String()
}
