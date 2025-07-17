// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package menu represents a DaisyUI Menu component.
//
// https://daisyui.com/components/menu/
package menu

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

// Props represents the properties for a DaisyUI menu component.
type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
}

// Option is a functional option to apply menu properties.
type Option components.Option[*Props]

func WithID(id string) Option {
	return func(p *Props) {
		p.attributes.SetID(id)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

// WithExtraClasses assigns additional classes to the Component.
func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// Build generates menu properties with the given options.
func Build(options ...Option) *Props {
	menu := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(menu)
	}

	return menu
}
