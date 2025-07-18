// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package hiddeninput contains methods for generating hidden inputs.
package hiddeninput

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/input"
)

type Option components.Option[*Props]

type Props struct {
	attributes *attributes.Attributes
}

// WithName will set the name attribute of the input.
func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.attributes.SetName(name)
	}
}

// WithID will set the id attribute of the input.
func WithID(id string) Option {
	return func(p *Props) {
		p.attributes.SetID(id)
	}
}

// WithValue will set the value attribute of the input.
func WithValue(value attributes.Value) Option {
	return func(p *Props) {
		p.attributes.SetValue(value)
	}
}

// WithExtraAttributes will set additional attributes on the input.
func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

func Build(options ...Option) *Props {
	hiddeninput := &Props{
		attributes: attributes.New(),
	}

	hiddeninput.attributes.SetAttribute(input.AttrType, "hidden")

	for _, option := range options {
		option(hiddeninput)
	}

	return hiddeninput
}
