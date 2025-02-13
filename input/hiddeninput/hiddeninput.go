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
	*attributes.Attributes
}

// WithName will set the name attribute of the input.
func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
	}
}

// WithID will set the id attribute of the input.
func WithID(id attributes.ID) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

// WithValue will set the value attribute of the input.
func WithValue(value attributes.Value) Option {
	return func(p *Props) {
		p.SetValue(value)
	}
}

// WithExtraAttributes will set additional attributes on the input.
func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func Build(options ...Option) *Props {
	hiddeninput := &Props{
		Attributes: attributes.New(),
	}

	hiddeninput.SetAttribute(input.AttrType, "hidden")

	for _, option := range options {
		option(hiddeninput)
	}

	return hiddeninput
}
