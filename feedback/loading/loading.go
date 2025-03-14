// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package loading represents a DaisyUI loading style.
//
// https://daisyui.com/components/loading/
package loading

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

const (
	Spinner LoadingStyle = iota
	Dots
	Ring
	Ball
	Bars
	Infinity
)

type LoadingStyle int

type Option components.Option[*Props]

// Props are the properties of a Loading component.
type Props struct {
	style         LoadingStyle
	htmxIndicator bool
	*attributes.Attributes
	size size.ResponsiveSize
	*color.ThemeColorProps
	*color.StateColorProps
}

// AsHTMXIndicator adds the "htmx-indicator" class to the Loading component,
// which will allow HTMX to add styling and control to the component.
func AsHTMXIndicator() Option {
	return func(lp *Props) {
		lp.htmxIndicator = true
	}
}

func WithID(id string) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
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

// Build creates a new LoadingProps.
func Build(style LoadingStyle, options ...Option) *Props {
	loading := &Props{
		style:           style,
		Attributes:      attributes.New(),
		ThemeColorProps: &color.ThemeColorProps{},
		StateColorProps: &color.StateColorProps{},
	}

	for _, option := range options {
		option(loading)
	}

	return loading
}
