// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package link

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

type Option components.Option[*Props]

type Props struct {
	attributes       *attributes.Attributes
	classes          *components.Classes
	underlineOnHover bool
}

func WithColor(color components.Class) Option {
	return func(p *Props) {
		p.classes.Add(color)
	}
}

func WithHref(href string) Option {
	return func(p *Props) {
		p.attributes.SetAttribute("href", href)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		if attrs != nil {
			p.attributes.AddAttributes(attrs)
		}
	}
}

// WithExtraClasses sets additional CSS classes on the component.
func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

func Build(options ...Option) *Props {
	link := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(link)
	}

	return link
}
