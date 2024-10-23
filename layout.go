// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate stringer -type=Layout -linecomment -output layout_generated.go
package main

// Layout options. Vertical is the default in DaisyUI so it goes first.
const (
	VerticalLayout   Layout = iota // vertical
	HorizontalLayout               // horizontal
)

type Layout int

func AddLayout(object string, layout Layout) string {
	return object + "-" + layout.String()
}
