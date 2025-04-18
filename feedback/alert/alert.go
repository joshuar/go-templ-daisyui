// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Alert is a DaisyUI alert component.
//
// https://daisyui.com/components/alert/
package alert

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

type Option components.Option[*Props]

type Props struct {
	*attributes.Attributes
	state AlertState
	style AlertStyle
	text  string
}

func WithText(text string) Option {
	return func(p *Props) {
		p.text = text
	}
}

func WithID(id string) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

// WithStyle will apply the given style to the component.
func WithStyle(style AlertStyle) Option {
	return func(p *Props) {
		p.style = style
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// Build creates a new Alert without rendering it. The Alert properties can then
// be modified before finally rendering by calling the Show() method.
func Build(state AlertState, options ...Option) *Props {
	alert := &Props{
		state:      state,
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(alert)
	}

	return alert
}
