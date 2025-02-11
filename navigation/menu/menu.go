// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package menu

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/breakpoints"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

const (
	Vertical Layout = iota
	Horizontal
)

type Layout int

type (
	Option components.Option2[*Props]
	Props  struct {
		title  string
		layout Layout
		*attributes.Attributes
		size      size.ResponsiveSize
		baseColor color.BaseColor
		breakpoints.Breakpoints
		// components.componentItems
		// componentHiddenBreakpoint
		// componentRevealedBreakpoint
		// modifierResponsiveSize
		// modifierLayout
		// modifierBaseColor
	}
)

// WithMenuTitle sets the menu title.
func WithMenuTitle(title string) Option {
	return func(p *Props) {
		p.title = title
	}
}

// WithLayout sets the layout of the menu, either vertically or horizontally. If
// this option is not specified, the menu will default to a vertical layout.
func WithLayout(layout Layout) Option {
	return func(p *Props) {
		p.layout = layout
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

func WithSize(menuSize size.ResponsiveSize) Option {
	return func(p *Props) {
		p.size = menuSize
	}
}

func WithBaseColor(base color.BaseColor) Option {
	return func(p *Props) {
		p.baseColor = base
	}
}

// NewMenu creates a new menu with the given options.
func Build(options ...Option) *Props {
	menu := &Props{
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(menu)
	}

	return menu
}

// type DropDownMenu struct {
// 	Props
// 	modifierBaseColor
// }

// // NewDropDownMenu creates a new menu wrapped in a dropdown with the given options.
// func NewDropDownMenu(options ...Option[*DropDownMenu]) *DropDownMenu {
// 	menu := &DropDownMenu{}

// 	for _, option := range options {
// 		menu = option(menu)
// 	}

// 	return menu
// }
