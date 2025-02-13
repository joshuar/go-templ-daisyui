// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package text provides objects and methods to render text enclosed in <p></p>
// tags, with the specified Tailwind text classes applied.
package text

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

const (
	Base Size = iota
	SM
	LG
	XL
	XLXL
)

const (
	Normal Weight = iota
	Thin
	Extralight
	Light
	Medium
	Semibold
	Bold
	Extrabold
	Black
)

type (
	// Text Size
	//
	// https://tailwindcss.com/docs/font-size
	Size int
	// Text Weight
	//
	// https://tailwindcss.com/docs/font-weight
	Weight int
)

type Option components.Option[*Props]

type Props struct {
	size   Size
	weight Weight
	italic bool
	text   string
	*color.ThemeColorProps
	*color.StateColorProps
}

// WithTextSize option sets the font size.
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

func Build(text string, options ...Option) *Props {
	props := &Props{
		text: text,
	}

	for _, option := range options {
		option(props)
	}

	return props
}
