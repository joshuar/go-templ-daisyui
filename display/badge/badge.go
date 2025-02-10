// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Badge is a DaisyUI Badge component.
//
// https://daisyui.com/components/badge/
package badge

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

type Option components.Option2[*Props]

// BadgeProps represents the properties for a Badge.
type Props struct {
	size size.ResponsiveSize
	color.Colors
	text string
}

// WithText will set the text to display within the Badge.
func WithText(text string) Option {
	return func(p *Props) {
		p.text = text
	}
}

// WithText will set the text to display within the Badge.
func WithSize(badgeSize size.ResponsiveSize) Option {
	return func(p *Props) {
		p.size = badgeSize
	}
}

func WithThemeColor(themeColor color.ThemeColor, outline bool) Option {
	return func(p *Props) {
		p.SetColor(themeColor, outline)
	}
}

func WithStateColor(stateColor color.StateColor, outline bool) Option {
	return func(p *Props) {
		p.SetState(stateColor, outline)
	}
}

// Build creates a new Badge without rendering it. The Badge properties can then
// be modified before finally rendering by calling the Show() method.
func Build(options ...Option) *Props {
	badge := &Props{}

	for _, option := range options {
		option(badge)
	}

	return badge
}
