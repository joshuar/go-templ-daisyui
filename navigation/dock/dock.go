// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package dock

import (
	"slices"

	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/actions/button"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
	buttons    []*button.Props
}

func Build(options ...Option) *Props {
	dock := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for option := range slices.Values(options) {
		option(dock)
	}

	return dock
}

type Option components.Option[*Props]

func WithButton(icon templ.Component, label string, attributes templ.Attributes) Option {
	return func(p *Props) {
		p.buttons = append(p.buttons, button.Build(
			button.WithContent(dockButtonContent(icon, label)),
			button.WithExtraAttributes(attributes),
		))
	}
}

func WithSize(size Size) Option {
	return func(p *Props) {
		p.classes.Add(size)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

func WithID(id string) Option {
	return func(p *Props) {
		p.attributes.SetID(id)
	}
}
