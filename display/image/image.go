// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package image

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/layout/mask"
)

type Option components.Option[*Props]

// Props represents the properties for an image.
type Props struct {
	url        string
	alt        string
	classes    *components.Classes
	attributes *attributes.Attributes
}

// WithLazyLoading will add the `loading="lazy"` attribute to the image.
func WithLazyLoading() Option {
	return func(image *Props) {
		image.attributes.SetAttribute("loading", "lazy")
	}
}

// WithAltText sets the alt text to be displayed for the image.
func WithAltText(alt string) Option {
	return func(image *Props) {
		image.alt = alt
	}
}

func WithMask(imageMask mask.Mask) Option {
	return func(image *Props) {
		image.classes.Add(imageMask)
	}
}

func WithExtraClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// Build creates a new image with the given URL and options.
func Build(url string, options ...Option) *Props {
	image := &Props{
		url:        url,
		classes:    components.NewClasses(),
		attributes: attributes.New(),
	}

	for _, option := range options {
		option(image)
	}

	return image
}
