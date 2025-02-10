// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package size contains modifiers for setting relative/responsive sizes on
// Components.
package size

const (
	MinResponsive ResponsiveSize = iota
	XSXS                         // Extra-extra small size. class: {component}-2xs
	XS                           // Extra-small size. class: {component}-xs
	SM                           // Small size. class: {component}-sm
	MD                           // Medium size. class: {component}-md
	LG                           // Large size. class: {component}-lg
	XL                           // Extra-large size. class: {component}-xl
	XLXL                         // Extra-extra large size. class: {component}-2xl
	MaxResponsive
)

// ResponsiveSize represents a relative/responsive size.
type ResponsiveSize int

// Valid returns true if the size is a valid response size.
func (size ResponsiveSize) Valid() bool {
	return size > MinResponsive && size < MaxResponsive
}
