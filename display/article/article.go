// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package article

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
)

const (
	Gray    ProseGrayScale = iota + 1 // prose-gray
	Slate                             // prose-slate
	Zinc                              // prose-zinc
	Neutral                           // prose-neutral
	Stone                             // prose-stone
)

type ProseGrayScale int

type Option components.Option[*Props]

type Props struct {
	grayscale ProseGrayScale
	content   templ.Component
	*attributes.Attributes
	*classes.Classes
}

// WithGrayScale adds optional gray scale scheme to the prose.
// https://github.com/tailwindlabs/tailwindcss-typography/blob/main/README.md#choosing-a-gray-scale
func WithGrayScale(scale ProseGrayScale) Option {
	return func(a *Props) {
		a.grayscale = scale
	}
}

// WithExtraAttributes assigns additional attributes to the Component.
func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// WithExtraClasses assigns additional classes to the Component.
func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		p.AddClasses(extraClasses...)
	}
}

// BuildArticle creates a new article component with the given title, content and
// options. Title is optional.
func Build(content templ.Component, options ...Option) *Props {
	article := &Props{
		content:    content,
		Attributes: attributes.New(),
		Classes:    classes.New(),
	}

	for _, option := range options {
		option(article)
	}

	return article
}
