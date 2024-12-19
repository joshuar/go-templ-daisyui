// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

type Image struct {
	url string
	alt string
	componentClasses
}

// WithURL sets the URL for the image.
func WithURL(url string) Option[Image] {
	return func(i Image) Image {
		i.url = url
		return i
	}
}

// WithAltText sets the alt text to be displayed for the image.
func WithAltText(alt string) Option[Image] {
	return func(i Image) Image {
		i.alt = alt
		return i
	}
}

// WithMask will apply a mask over the image.
func WithMask(mask Mask) Option[Image] {
	return func(i Image) Image {
		i.AddClasses("mask", mask.String())
		return i
	}
}

// WithObjectFit will add an Object Fit setting to the image class.
func WithImageObjectFit(fit ObjectFit) Option[Image] {
	return func(i Image) Image {
		if fit > 0 {
			i.AddClasses(fit.String())
		}

		return i
	}
}

// NewImage creates a new image component with the given options. The created
// image can be rendered by calling its Show method.
func NewImage(options ...Option[Image]) Image {
	image := Image{}

	for _, option := range options {
		image = option(image)
	}

	return image
}
