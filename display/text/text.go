// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package text provides objects and methods to render text enclosed in <p></p>
// tags, with the specified Tailwind text classes applied.
package text

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

type Option components.Option[*Props]

type Props struct {
	size   Size
	weight Weight
	italic bool
	text   string
	*color.ThemeColorProps
	*color.StateColorProps
	*attributes.Attributes
}

// WithTextSize option sets the font size. If this option is not specified,
// 'text-base' is used.
//
// https://tailwindcss.com/docs/font-size
func WithTextSize(size Size) Option {
	return func(text *Props) {
		text.size = size
	}
}

// WithTextWeight option sets the font weight.
//
// https://tailwindcss.com/docs/font-weight
func WithTextWeight(weight Weight) Option {
	return func(text *Props) {
		text.weight = weight
	}
}

// AsItalicText option will render the text as italic.
//
// https://tailwindcss.com/docs/font-style
func AsItalicText() Option {
	return func(text *Props) {
		text.italic = true
	}
}

func WithThemeColor(themeColor color.ThemeColor) Option {
	return func(p *Props) {
		p.ThemeColorProps = color.NewThemeColor(themeColor, false)
	}
}

func WithStateColor(stateColor color.StateColor) Option {
	return func(p *Props) {
		p.StateColorProps = color.NewStateColor(stateColor, false)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func Build(text string, options ...Option) *Props {
	props := &Props{
		text:       text,
		size:       Base,
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(props)
	}

	return props
}
