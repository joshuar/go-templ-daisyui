// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Toggle represents a DaisyUI Toggle Component.
//
// https://daisyui.com/components/toggle/
package toggle

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

type Option components.Option[*Props]

type Props struct {
	*attributes.Attributes
	size size.ResponsiveSize
	*color.ThemeColorProps
	*color.StateColorProps
	disabled bool
}

func WithChecked(value bool) Option {
	return func(p *Props) {
		if value {
			p.Checked()
		}

		p.UnChecked()
	}
}

func WithSize(toggleSize size.ResponsiveSize) Option {
	return func(p *Props) {
		p.size = toggleSize
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

func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
	}
}

func WithID(id attributes.ID) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithValue(value attributes.Value) Option {
	return func(p *Props) {
		p.SetValue(value)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func Build(options ...Option) *Props {
	toggle := &Props{
		Attributes:      attributes.New(),
		ThemeColorProps: &color.ThemeColorProps{},
		StateColorProps: &color.StateColorProps{},
	}

	for _, option := range options {
		option(toggle)
	}

	return toggle
}
