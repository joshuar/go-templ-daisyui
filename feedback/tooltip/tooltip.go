// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package tooltip represents a DaisyUI tooltip component.
//
// https://daisyui.com/components/tooltip/
package tooltip

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

// Option is an option to apply to a tooltip.
type Option components.Option[*Props]

// Props defines the properties of the tooltip.
type Props struct {
	classes    *components.Classes
	attributes *attributes.Attributes
	tip        string
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(tt *Props) {
		tt.attributes.AddAttributes(attrs)
	}
}

func WithClasses(classes ...components.Class) Option {
	return func(tt *Props) {
		tt.classes.Add(classes...)
	}
}

// Build generates tooltip properties with the given options.
func Build(tip string, options ...Option) *Props {
	tooltip := &Props{
		tip:        tip,
		classes:    components.NewClasses(),
		attributes: attributes.New(),
	}

	for _, option := range options {
		option(tooltip)
	}

	return tooltip
}
