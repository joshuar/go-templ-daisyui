// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package fieldset

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/input"
)

// FieldProps represents an input field in a fieldset.
type FieldProps struct {
	label string
	input input.Input
}

// Props represents a fieldset.
type Props struct {
	legend     string
	attributes *attributes.Attributes
	classes    *components.Classes
}

type Option components.Option[*Props]

// WithLegend option sets the legend (title) on the fieldset.
func WithLegend(legend string) Option {
	return func(fieldset *Props) {
		fieldset.legend = legend
	}
}

// WithID option sets the id attribute on the component.
func WithID(id string) Option {
	return func(fieldset *Props) {
		fieldset.attributes.SetID(id)
	}
}

// WithExtraAttributes option sets additional attributes on the component.
func WithAttributes(attrs templ.Attributes) Option {
	return func(fieldset *Props) {
		fieldset.attributes.AddAttributes(attrs)
	}
}

// WithExtraClasses option assigns additional classes to the component.
func WithClasses(extraClasses ...components.Class) Option {
	return func(fieldset *Props) {
		fieldset.classes.Add(extraClasses...)
	}
}

func Build(options ...Option) *Props {
	props := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(props)
	}

	return props
}
