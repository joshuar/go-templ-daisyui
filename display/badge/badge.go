// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package badge is a DaisyUI Badge component.
//
// https://daisyui.com/components/badge/
package badge

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

type Option components.Option[*Props]

// Props represents the properties for a Badge.
type Props struct {
	content    templ.Component
	classes    *components.Classes
	attributes *attributes.Attributes
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

// Build creates a new Badge without rendering it. The Badge properties can then
// be modified before finally rendering by calling the Show() method.
func Build(options ...Option) *Props {
	badge := &Props{
		classes:    components.NewClasses(),
		attributes: attributes.New(),
	}

	for _, option := range options {
		option(badge)
	}

	return badge
}
