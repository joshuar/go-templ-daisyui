// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package style provides modifiers for styling DaisyUI components.
package style

const (
	Min Style = iota
	Soft
	Outline
	Dash
	Max
)

// Style is the DaisyUI style to apply.
type Style int

// Valid returns true if the style is a valid style value.
func (style Style) Valid() bool {
	return style > Min && style < Max
}
