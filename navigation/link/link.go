// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package link

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

type Option components.Option[*Props]

type Props struct {
	*attributes.Attributes
	*classes.Classes
	stateColor       color.StateColor
	themeColor       color.ThemeColor
	content          templ.Component
	underlineOnHover bool
}

func WithThemeColor(themeColor color.ThemeColor) Option {
	return func(p *Props) {
		p.themeColor = themeColor
	}
}

func WithStateColor(stateColor color.StateColor) Option {
	return func(p *Props) {
		p.stateColor = stateColor
	}
}

func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
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

// WithExtraClasses sets additional CSS classes on the component.
func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		p.AddClasses(extraClasses...)
	}
}

// WithContent sets the content for the Button.
func WithContent(content any) Option {
	return func(p *Props) {
		p.content = components.SetContent(content)
	}
}

func WithUnderlineOnHover() Option {
	return func(p *Props) {
		p.underlineOnHover = true
	}
}

func Build(options ...Option) *Props {
	link := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
	}

	for _, option := range options {
		option(link)
	}

	return link
}
