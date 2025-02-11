// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package navbar represents a DaisyUI NavBar Component.
//
// https://daisyui.com/components/navbar/
package navbar

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

type Option components.Option2[*Props]

type Props struct {
	*attributes.Attributes
	baseColor color.BaseColor
	// modifierZIndex
	// modifierPosition
	// modifierBaseColor
	// HtmlAttrID
	start  templ.Component
	center templ.Component
	end    templ.Component
}

func NavBarStart(component templ.Component) Option {
	return func(navbar *Props) {
		navbar.start = component
	}
}

func NavBarCenter(component templ.Component) Option {
	return func(navbar *Props) {
		navbar.center = component
	}
}

func NavBarEnd(component templ.Component) Option {
	return func(navbar *Props) {
		navbar.end = component
	}
}

func WithID(id attributes.ID) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func WithBaseColor(base color.BaseColor) Option {
	return func(p *Props) {
		p.baseColor = base
	}
}

func Build(options ...Option) *Props {
	navbar := &Props{
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(navbar)
	}

	return navbar
}
