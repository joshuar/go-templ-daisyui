// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Icon is a Font Awesome (free) icon.
package icon

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

const (
	Solid   FontAwesomeStyle = iota // fa-solid
	Regular                         // fa-regular
)

type Option components.Option2[*Props]

// FontAwesomeStyle is a valid FontAwesome Icon style.
type FontAwesomeStyle int

// Props are the properties that can be set on an icon.
type Props struct {
	name         string
	style        FontAwesomeStyle
	relativeSize size.ResponsiveSize
	*color.ThemeColorProps
	*color.StateColorProps
}

// WithStyle assigns the given FontAwesome style to the Icon.
func WithStyle(style FontAwesomeStyle) Option {
	return func(p *Props) {
		p.style = style
	}
}

// WithRelativeSize sets a relative/responsive size for the icon.
//
// https://docs.fontawesome.com/web/style/size
func WithRelativeSize(relativeSize size.ResponsiveSize) Option {
	return func(p *Props) {
		p.relativeSize = relativeSize
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

// New creates iconProps with the given options.
func Build(name string, options ...Option) *Props {
	icon := &Props{
		name:            name,
		ThemeColorProps: &color.ThemeColorProps{},
		StateColorProps: &color.StateColorProps{},
	}

	for _, option := range options {
		option(icon)
	}

	return icon
}
