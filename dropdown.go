// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type DropdownProps struct {
	ModifierOpenFrom
	openOnHover bool
	alignToEnd  bool
}

// WithOpenOnHover option will ensure the Dropdown will open on hover.
func WithOpenOnHover() Option[*DropdownProps] {
	return func(dd *DropdownProps) *DropdownProps {
		dd.openOnHover = true
		return dd
	}
}

// WithAlignToEnd option aligns the dropdown content to the end of the container.
func WithAlignToEnd() Option[*DropdownProps] {
	return func(dd *DropdownProps) *DropdownProps {
		dd.alignToEnd = true
		return dd
	}
}

// FromDropdownProps option can be used to pass a pre-build DropdownProps to the
// Dropdown Component to render.
func FromDropdownProps(props *DropdownProps) Option[*DropdownProps] {
	return func(_ *DropdownProps) *DropdownProps {
		return props
	}
}

// BuildDropdown builds a DropdownProps from the given options. Use this to
// pre-build a Dropdown Component before rendering.
func BuildDropdown(options ...Option[*DropdownProps]) *DropdownProps {
	dropdown := &DropdownProps{}

	for _, option := range options {
		dropdown = option(dropdown)
	}

	return dropdown
}
