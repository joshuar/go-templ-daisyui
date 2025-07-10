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
)

type Option components.Option[*Props]

type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
}

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

func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

func Build(options ...Option) *Props {
	navbar := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(navbar)
	}

	return navbar
}
