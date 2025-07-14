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

// Option is a functional option to set text properties.
type Option components.Option[*Props]

// Props represents the properties of the text.
type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
}

// WithSize option sets the font size. If this option is not specified,
// 'text-base' is used.
func WithSize(size Size) Option {
	return func(p *Props) {
		p.classes.Add(size)
	}
}

// WithWeight option sets the font weight.
func WithWeight(weight Weight) Option {
	return func(text *Props) {
		text.classes.Add(weight)
	}
}

// AsItalicText option will render the text as italic.
func AsItalicText() Option {
	return func(text *Props) {
		text.classes.Add(Italic)
	}
}

// WithColor will add a DaisUI color class to the text. This sets the "content" color.
func WithColor(color Color) Option {
	return func(text *Props) {
		text.classes.Add(color)
	}
}

// WithOverflow option adds the given overflow class to the text.
func WithOverflow(flow Overflow) Option {
	return func(text *Props) {
		text.classes.Add(flow)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

func WithExtraClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
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
