// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package menu

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/items"
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
	Option components.Option[*Props]
	Props  struct {
		title  string
		layout Layout
		*attributes.Attributes
		*classes.Classes
		size      size.ResponsiveSize
		baseColor color.BaseColor
		breakpoints.Breakpoints
		*items.Items
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

// WithItems option sets the list of items to display in the dropdown menu. Each
// item will be wrapped in an list element.
func WithItems(listItems ...templ.Component) Option {
	return func(p *Props) {
		p.ReplaceItems(listItems...)
	}
}

func WithID(id string) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// WithExtraClasses assigns additional classes to the Component.
func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		for _, class := range extraClasses {
			p.AddClass(class)
		}
	}
}

// NewMenu creates a new menu with the given options.
func Build(options ...Option) *Props {
	menu := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
		Items:      items.New(),
	}

	for _, option := range options {
		option(menu)
	}

	return menu
}
