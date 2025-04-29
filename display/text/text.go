// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package text provides objects and methods to render text enclosed in <p></p>
// tags, with the specified Tailwind text classes applied.
package text

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

type Option components.Option[*Props]

type Props struct {
	text any
	*attributes.Attributes
	classes *components.Classes
}

// WithTextSize option sets the font size. If this option is not specified,
// 'text-base' is used.
//
// https://tailwindcss.com/docs/font-size
func WithTextSize(size Size) Option {
	return func(text *Props) {
		text.classes.Add(size)
	}
}

// WithTextWeight option sets the font weight.
//
// https://tailwindcss.com/docs/font-weight
func WithTextWeight(weight Weight) Option {
	return func(text *Props) {
		text.classes.Add(weight)
	}
}

// AsItalicText option will render the text as italic.
//
// https://tailwindcss.com/docs/font-style
func AsItalicText() Option {
	return func(text *Props) {
		text.classes.Add(Italic)
	}
}

// WithTextColor will add a DaisUI color class to the text. This sets the "content" color.
func WithTextColor(color Color) Option {
	return func(text *Props) {
		text.classes.Add(color)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func WithExtraClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

func Build(text any, options ...Option) *Props {
	props := &Props{
		text:       text,
		Attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(props)
	}

	return props
}
