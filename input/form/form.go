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
	*attributes.Attributes
	Elements []templ.Component
}

// WithID will set the id attribute on the form.
func WithID(id attributes.ID) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

// WithExtraAttributes will add the given attributes to the form.
func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// WithElements are the elements/components to show in the form.
func WithElements(elements ...templ.Component) Option {
	return func(p *Props) {
		p.Elements = elements
	}
}

// Build creates a new Form without rendering it. The Form properties can then
// be modified before finally rendering by calling the Show() method.
func BuildForm(options ...Option) *Props {
	alert := &Props{
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(alert)
	}

	return alert
}
