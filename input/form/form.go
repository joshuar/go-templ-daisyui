// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package form represents components for displaying HTML forms.
package form

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

type Option components.Option[*Props]

type Props struct {
	attributes *attributes.Attributes
}

// WithID will set the id attribute on the form.
func WithID(id string) Option {
	return func(p *Props) {
		p.attributes.SetID(id)
	}
}

// WithExtraAttributes will add the given attributes to the form.
func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

// Build creates a new Form without rendering it. The Form properties can then
// be modified before finally rendering by calling the Show() method.
func Build(options ...Option) *Props {
	alert := &Props{
		attributes: attributes.New(),
	}

	for _, option := range options {
		option(alert)
	}

	return alert
}
