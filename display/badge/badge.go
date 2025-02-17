// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Badge is a DaisyUI Badge component.
//
// https://daisyui.com/components/badge/
package badge

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

type Option components.Option[*Props]

// BadgeProps represents the properties for a Badge.
type Props struct {
	size size.ResponsiveSize
	*color.ThemeColorProps
	*color.StateColorProps
	content templ.Component
	*attributes.Attributes
}

// WithText will set the text to display within the Badge.
func WithContent(content any) Option {
	return func(p *Props) {
		p.content = components.SetContent(content)
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
		p.ThemeColorProps = color.NewThemeColor(themeColor, outline)
	}
}

func WithStateColor(stateColor color.StateColor, outline bool) Option {
	return func(p *Props) {
		p.StateColorProps = color.NewStateColor(stateColor, outline)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// Build creates a new Badge without rendering it. The Badge properties can then
// be modified before finally rendering by calling the Show() method.
func Build(options ...Option) *Props {
	badge := &Props{
		ThemeColorProps: &color.ThemeColorProps{},
		StateColorProps: &color.StateColorProps{},
		Attributes:      attributes.New(),
	}

	for _, option := range options {
		option(badge)
	}

	return badge
}
