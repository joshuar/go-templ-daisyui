// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package breadcrumbs

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/items"
)

type Option components.Option[*Props]

type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
	*items.Items
}

func WithID(id string) Option {
	return func(p *Props) {
		p.attributes.SetID(id)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

func WithClasses(classes ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(classes...)
	}
}

func WithCrumbs(crumbs ...templ.Component) Option {
	return func(p *Props) {
		p.ReplaceItems(crumbs...)
	}
}

func Build(options ...Option) *Props {
	link := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
		Items:      items.New(),
	}

	for _, option := range options {
		option(link)
	}

	return link
}
