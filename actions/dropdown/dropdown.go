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
	"github.com/joshuar/go-templ-daisyui/display/card"
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

// Open defines where the Dropdown will open from.
type Open int

const (
	// Details uses a <details> element for the Dropdown.
	//
	// https://daisyui.com/components/dropdown/#dropdown-menu-using-details-tag
	Details Style = iota
	// Class uses a <div> element for the Dropdown.
	//
	// https://daisyui.com/components/dropdown/#dropdown-menu
	Class
)

// Style defines what type of Dropdown to use. This can be either Details (using
// a <details> element) or Class (using a <div> element).
type Style int

// Content is an interface to represent any type of Component that can be set as
// the Dropdown content.
type Content interface {
	Show(classes ...string) templ.Component
}

// Props holds common properties for all types of Dropdown components.
type Props struct {
	style     Style
	openProps *OpenProps
	button    *button.Props
	*breakpoints.Breakpoints
	content Content
}

// Option is a functional option to apply to a Dropdown.
type Option components.Option[*Props]

// OpenProps defines properties for how a Dropdown will open.
type OpenProps struct {
	From         Open
	EndAlignment bool
	Hover        bool
	Force        bool
}

// OpenOption is a functional option to control how a Dropdown will open.
type OpenOption components.Option[*OpenProps]

// OnHover option will ensure the Dropdown will open on hover.
func OnHover() OpenOption {
	return func(p *OpenProps) {
		p.Hover = true
	}
}

// ForceOpen option will ensure the Dropdown will be open by default.
func ForceOpen() OpenOption {
	return func(p *OpenProps) {
		p.Force = true
	}
}

// From customizes how the dropdown will open and optional alignment. By default, if this
// option is not specified, the dropdown will open from the bottom.
func From(from Open, alignend bool) OpenOption {
	return func(p *OpenProps) {
		p.From = from
		p.EndAlignment = alignend
	}
}

// WithOpenOptions option sets the options for how the Dropdown will open.
func WithOpenOptions(options ...OpenOption) Option {
	return func(p *Props) {
		for _, option := range options {
			option(p.openProps)
		}
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
func WithMenuContent(options ...menu.Option) Option {
	return func(p *Props) {
		menu := menu.Build(options...)
		menu.SetAttribute("tabindex", 0)
		p.content = menu
	}
}

// WithCardContent option sets the options for a Card component as the Dropdown
// content.. Use this option to set the Card content and optional styling.
func WithCardContent(options ...card.Option) Option {
	return func(p *Props) {
		card := card.Build(options...)
		p.content = card
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

// AsStyle option sets the style of the Dropdown component.
func AsStyle(style Style) Option {
	return func(p *Props) {
		p.style = style
	}
}

// BuildDropdown builds a DropdownProps from the given options. Use this to
// pre-build a Dropdown Component before rendering.
func Build(options ...Option) *Props {
	dropdown := &Props{
		openProps:   &OpenProps{},
		Breakpoints: &breakpoints.Breakpoints{},
	}

	for _, option := range options {
		option(dropdown)
	}

	return dropdown
}
