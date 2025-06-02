// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package figure

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/display/image"
)

type Props struct {
	image      *image.Props
	classes    *components.Classes
	attributes *attributes.Attributes
}

type Option components.Option[*Props]

// Build creates a new image with the given URL and options.
func Build(img *image.Props, options ...Option) *Props {
	figure := &Props{
		image:      img,
		classes:    components.NewClasses(),
		attributes: attributes.New(),
	}

	for _, option := range options {
		option(figure)
	}

	return figure
}

func WithFigureAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

func WithFigureClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}
