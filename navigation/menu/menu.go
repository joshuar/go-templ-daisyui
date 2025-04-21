// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package menu represents a DaisyUI Menu component.
//
// https://daisyui.com/components/menu/
package menu

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/items"
	"github.com/joshuar/go-templ-daisyui/modifiers/breakpoints"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

// Props represents the properties for a DaisyUI menu component.
type Props struct {
	*components.Props
	title     templ.Component
	size      Size
	baseColor color.BaseColor
	breakpoints.Breakpoints
	*items.Items
	itemClasses *classes.Classes
}

// Option is a functional option to apply menu properties.
type Option components.Option[*Props]

// WithMenuTitle sets the menu title.
func WithMenuTitle(title any) Option {
	return func(p *Props) {
		p.title = components.SetContent(title)
	}
}

// WithLayout sets the layout of the menu, either vertically or horizontally. If
// this option is not specified, the menu will default to a vertical layout.
func WithLayout(layout Layout) Option {
	return func(p *Props) {
		p.Props.AddClasses(layout)
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

func WithSize(menuSize Size) Option {
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

// WithItemExtraClasses assigns additional classes to each item of the menu.
func WithExtraItemClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		p.Items.SetClasses(extraClasses...)
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
		p.Props.AddClasses(extraClasses...)
	}
}

// Build generates menu properties with the given options.
func Build(options ...Option) *Props {
	menu := &Props{
		Items:       items.New(),
		Props:       components.InitProps(),
		itemClasses: classes.New(),
	}

	for _, option := range options {
		option(menu)
	}

	return menu
}
