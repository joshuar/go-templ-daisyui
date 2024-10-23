// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import "github.com/a-h/templ"

type Menu struct {
	Layout Layout
	Size   Size
	Items  []MenuItem
}

type MenuItem struct {
	Tooltip string
	Label   string
	Content any
	Attrs   templ.Attributes
	State
}
