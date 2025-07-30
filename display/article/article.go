// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package article represents an article component, using Tailwind's typography plugin
// and "prose" class.
//
// https://daisyui.com/docs/layout-and-typography/#-1
package article

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

// Option is a functional option to set a property of an article.
type Option components.Option[*Props]

// Props are the properties of an article component.
type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
}

// WithAttributes assigns additional attributes to the Component.
func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

// WithClasses assigns additional classes to the Component.
func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// Build creates a new article component with the given options.
func Build(options ...Option) *Props {
	article := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(article)
	}

	return article
}
