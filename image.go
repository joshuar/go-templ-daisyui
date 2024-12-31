// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

type Image struct {
	url string
	alt string
	modifierMask
	modifierObjectFit
	componentClasses
}

// WithAltText sets the alt text to be displayed for the image.
func WithAltText(alt string) Option[Image] {
	return func(i Image) Image {
		i.alt = alt
		return i
	}
}

// NewImage creates a new image component with the given options. The created
// image can be rendered by calling its Show method.
func NewImage(url string, options ...Option[Image]) Image {
	image := Image{
		url: url,
	}

	for _, option := range options {
		image = option(image)
	}

	return image
}
