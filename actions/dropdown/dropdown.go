// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package dropdown defines a DaisyUI Dropdown.
//
// https://daisyui.com/components/dropdown
package dropdown

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/actions/button"
	"github.com/joshuar/go-templ-daisyui/modifiers/breakpoints"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
	"github.com/joshuar/go-templ-daisyui/navigation/menu"
)

const (
	OpenTop Open = iota + 1
	OpenBottom
	OpenLeft
	OpenRight
)

type Open int

type Option components.Option[*Props]

type Props struct {
	open        Open
	alignToEnd  bool
	openOnHover bool
	forceOpen   bool
	button      *button.Props
	menu        *menu.Props
	*breakpoints.Breakpoints
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

// WithOpenFrom customizes how the dropdown will open and optional alignment. By default, if this
// option is not specified, the dropdown will open from the bottom.
func WithOpenFrom(from Open, alignToEnd bool) Option {
	return func(p *Props) {
		p.open = from
		p.alignToEnd = alignToEnd
	}
}

// WithButton option sets the options for the dropdown button. If not specified,
// a default button will be used with minimal styling and the text "Click" inside.
func WithButton(options ...button.Option) Option {
	return func(p *Props) {
		p.button = button.Build(options...)
		button.WithExtraAttributes(templ.Attributes{
			"tabindex": 0,
		})(p.button)
	}
}

// WithMenuOptions option sets the options for the dropdown menu. Use this
// option to set optional menu styling and to define the list of items for the
// dropdown menu.
func WithMenuOptions(options ...menu.Option) Option {
	return func(p *Props) {
		p.menu = menu.Build(options...)
		p.menu.SetAttribute("tabindex", 0)
	}
}

func WithHiddenBreakpoint(from size.ResponsiveSize) Option {
	return func(p *Props) {
		p.SetHiddenBreakpoint(from)
	}
}

func WithRevealedBreakpoint(from size.ResponsiveSize) Option {
	return func(p *Props) {
		p.SetRevealedBreakpoint(from)
	}
}

// BuildDropdown builds a DropdownProps from the given options. Use this to
// pre-build a Dropdown Component before rendering.
func Build(options ...Option) *Props {
	dropdown := &Props{
		Breakpoints: &breakpoints.Breakpoints{},
	}

	for _, option := range options {
		option(dropdown)
	}

	return dropdown
}
