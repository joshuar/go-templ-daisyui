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
	"github.com/joshuar/go-templ-daisyui/modifiers/state"
	"github.com/joshuar/go-templ-daisyui/modifiers/style"
)

type Option components.Option[*Props]

type Props struct {
	*attributes.Attributes
	state state.State
	style style.Style
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
func WithStyle(value style.Style) Option {
	return func(p *Props) {
		p.style = value
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// Build creates a new Alert without rendering it. The Alert properties can then
// be modified before finally rendering by calling the Show() method.
func Build(alertState state.State, options ...Option) *Props {
	alert := &Props{
		state:      alertState,
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(alert)
	}

	return alert
}
