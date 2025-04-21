// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package badge is a DaisyUI Badge component.
//
// https://daisyui.com/components/badge/
package badge

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
)

type Option components.Option[*Props]

// Props represents the properties for a Badge.
type Props struct {
	color   BadgeColor
	style   BadgeStyle
	size    BadgeSize
	content templ.Component
	*attributes.Attributes
	*classes.Classes
}

// WithText will set the text to display within the Badge.
func WithContent(content any) Option {
	return func(p *Props) {
		p.content = components.SetContent(content)
	}
}

// WithText will set the text to display within the Badge.
func WithSize(size BadgeSize) Option {
	return func(p *Props) {
		p.size = size
	}
}

func WithColor(color BadgeColor) Option {
	return func(p *Props) {
		p.color = color
	}
}

// WithStyle will set a style for a colored Badge. Styles are Soft, Outline or
// DashedOutline.
func WithStyle(style BadgeStyle) Option {
	return func(p *Props) {
		p.style = style
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		p.AddClasses(extraClasses...)
	}
}

// Build creates a new Badge without rendering it. The Badge properties can then
// be modified before finally rendering by calling the Show() method.
func Build(options ...Option) *Props {
	badge := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
	}

	for _, option := range options {
		option(badge)
	}

	return badge
}
