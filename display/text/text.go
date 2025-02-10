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
	Base TextSize = iota
	SM
	LG
	XL
	XLXL
)

const (
	Normal TextWeight = iota
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
	TextSize int
	// Text Weight
	//
	// https://tailwindcss.com/docs/font-weight
	TextWeight int
)

type Option components.Option2[*Props]

type Props struct {
	size   TextSize
	weight TextWeight
	italic bool
	color.Colors
	text string
}

// WithTextSize option sets the font size.
//
// https://tailwindcss.com/docs/font-size
func WithTextSize(size TextSize) Option {
	return func(text *Props) {
		text.size = size
	}
}

// WithTextWeight option sets the font weight.
//
// https://tailwindcss.com/docs/font-weight
func WithTextWeight(weight TextWeight) Option {
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

func Build(text string, options ...Option) *Props {
	props := &Props{
		text: text,
	}

	for _, option := range options {
		option(props)
	}

	return props
}
