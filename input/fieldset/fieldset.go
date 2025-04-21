// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package fieldset

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/input"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

// FieldProps represents an input field in a fieldset.
type FieldProps struct {
	label string
	input input.Input
}

// Props represents a fieldset.
type Props struct {
	legend    string
	fields    []*FieldProps
	baseColor color.BaseColor
	*attributes.Attributes
	*classes.Classes
}

type Option components.Option[*Props]

// WithLegend option sets the legend (title) on the fieldset.
func WithLegend(legend string) Option {
	return func(fieldset *Props) {
		fieldset.legend = legend
	}
}

// WithField option adds the given input with an optional label as a field in
// the fieldset. Call this multiple times to add multiple inputs.
func WithField(label string, field input.Input) Option {
	return func(fieldset *Props) {
		fieldset.fields = append(fieldset.fields, &FieldProps{label: label, input: field})
	}
}

// WithID option sets the id attribute on the component.
func WithID(id string) Option {
	return func(fieldset *Props) {
		fieldset.SetID(id)
	}
}

// WithExtraAttributes option sets additional attributes on the component.
func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(fieldset *Props) {
		fieldset.AddAttributes(attrs)
	}
}

// WithExtraClasses option assigns additional classes to the component.
func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(fieldset *Props) {
		fieldset.AddClasses(extraClasses...)
	}
}

// WithBaseColor option sets a base color on the component.
func WithBaseColor(baseColor color.BaseColor) Option {
	return func(fieldset *Props) {
		fieldset.baseColor = baseColor
	}
}

func Build(options ...Option) *Props {
	props := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
	}

	for _, option := range options {
		option(props)
	}

	return props
}
