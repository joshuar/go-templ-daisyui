// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package image

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/layout/mask"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

type Option components.Option[*Props]

// Props represents the properties for an image.
type Props struct {
	url         string
	alt         string
	lazyLoading bool
	mask        mask.Mask
	size        size.FixedSize
	*classes.Classes
	// modifierObjectFit
}

// WithLazyLoading will add the `loading="lazy"` attribute to the image.
func WithLazyLoading() Option {
	return func(image *Props) {
		image.lazyLoading = true
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
		image.mask = imageMask
	}
}

func WithSize(imageSize size.FixedSize) Option {
	return func(p *Props) {
		p.size = imageSize
	}
}

func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		for _, class := range extraClasses {
			p.AddClass(class)
		}
	}
}

// BuildImage creates a new ImageProps.
func Build(url string, options ...Option) *Props {
	image := &Props{
		url:     url,
		Classes: classes.New(),
	}

	for _, option := range options {
		option(image)
	}

	return image
}
