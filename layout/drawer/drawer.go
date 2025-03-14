// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package drawer

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

type Option components.Option[*Props]

type Props struct {
	*attributes.Attributes
	sideVisibility size.ResponsiveSize
	sideContent    templ.Component
	mainContent    templ.Component
}

// WithDrawerVisibility sets the size at which the drawer will always be
// visible. Below this value, the drawer will be hidden by default.
func WithDrawerVisibility(visbility size.ResponsiveSize) Option {
	return func(p *Props) {
		p.sideVisibility = visbility
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// NewCard creates a new card component with the given options. The card can be
// rendered by calling the Show method.
func Build(id string, main, side templ.Component, options ...Option) *Props {
	drawer := &Props{
		mainContent: main,
		sideContent: side,
		Attributes:  attributes.New(),
	}

	drawer.SetID(id)

	for _, option := range options {
		option(drawer)
	}

	return drawer
}
