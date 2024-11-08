// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Layout -linecomment -output layout_generated.go
package components

// Layout options. Vertical is the default in DaisyUI so it goes first.
const (
	VerticalLayout   Layout = iota // vertical
	HorizontalLayout               // horizontal
)

type Layout int

// LayoutObject returns the given object class with an appropriate layout
// modifier. For e.g., if Layout is VerticalLayout and object is menu, the return
// value will be "menu-vertical".
func (l Layout) LayoutObject(object string) string {
	return object + "-" + l.String()
}
