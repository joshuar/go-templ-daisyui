// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package tooltip represents a DaisyUI tooltip component.
//
// https://daisyui.com/components/tooltip/
package tooltip

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

// Option is an option to apply to a tooltip.
type Option components.Option[*Props]

// Props defines the properties of the tooltip.
type Props struct {
	from       OpenFrom
	stateColor color.StateColor
	themeColor color.ThemeColor
	text       string
}

// WithThemeColor option sets a theme color for the component.
func WithThemeColor(themeColor color.ThemeColor) Option {
	return func(p *Props) {
		p.themeColor = themeColor
	}
}

// WithStateColor option sets a state color for the component.
func WithStateColor(stateColor color.StateColor) Option {
	return func(p *Props) {
		p.stateColor = stateColor
	}
}

// WithOpenFrom option sets where the tooltip will open from. If not specified, the tooltip will open from above the
// element.
func WithOpenFrom(from OpenFrom) Option {
	return func(p *Props) {
		p.from = from
	}
}

// Build generates tooltip properties with the given options.
func Build(text string, options ...Option) *Props {
	tooltip := &Props{
		text: text,
	}

	for _, option := range options {
		option(tooltip)
	}

	return tooltip
}
