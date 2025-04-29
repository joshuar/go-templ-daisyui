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
	*attributes.Attributes
	classes *components.Classes
}

// WithGrayScale adds optional gray scale scheme to the prose.
// https://github.com/tailwindlabs/tailwindcss-typography/blob/main/README.md#choosing-a-gray-scale
func WithGrayScale(scale GrayScale) Option {
	return func(a *Props) {
		a.classes.Add(scale)
	}
}

// WithExtraAttributes assigns additional attributes to the Component.
func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// WithExtraClasses assigns additional classes to the Component.
func WithExtraClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// Build creates a new article component with the given options.
func Build(options ...Option) *Props {
	article := &Props{
		Attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(article)
	}

	return article
}
