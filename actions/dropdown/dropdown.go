// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package dropdown defines a DaisyUI Dropdown.
//
// https://daisyui.com/components/dropdown
package dropdown

import components "github.com/joshuar/go-templ-daisyui"

const (
	OpenTop Open = iota + 1
	OpenBottom
	OpenLeft
	OpenRight
)

type Open int

type Option components.Option2[*Props]

type Props struct {
	open        Open
	alignToEnd  bool
	openOnHover bool
	forceOpen   bool
}

// WithOpenOnHover option will ensure the Dropdown will open on hover.
func WithOpenOnHover() Option {
	return func(p *Props) {
		p.openOnHover = true
	}
}

// WithForceOpen will ensure the Dropdown will be open by default.
func WithForceOpen() Option {
	return func(p *Props) {
		p.forceOpen = true
	}
}

// WithOpenFrom customises how the dropdown will open and optional alignment. By default, if this
// option is not specified, the dropdown will open from the bottom.
func WithOpenFrom(from Open, alignToEnd bool) Option {
	return func(p *Props) {
		p.open = from
		p.alignToEnd = alignToEnd
	}
}

// BuildDropdown builds a DropdownProps from the given options. Use this to
// pre-build a Dropdown Component before rendering.
func Build(options ...Option) *Props {
	dropdown := &Props{}

	for _, option := range options {
		option(dropdown)
	}

	return dropdown
}
