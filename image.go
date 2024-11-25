// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "strings"

type Image struct {
	URL   string
	Alt   string
	class []string
}

func (i Image) Class() string {
	return strings.Join(i.class, " ")
}

// CardOption adds an option to a card.
type ImageOption Option[Image]

// WithURL sets the URL for the image.
func WithURL(url string) ImageOption {
	return func(i Image) Image {
		i.URL = url
		return i
	}
}

// WithAltText sets the alt text to be displayed for the image.
func WithAltText(alt string) ImageOption {
	return func(i Image) Image {
		i.Alt = alt
		return i
	}
}

// WithMask will apply a mask over the image.
func WithMask(mask Mask) ImageOption {
	return func(i Image) Image {
		i.class = append(i.class, "mask", mask.String())
		return i
	}
}

// WithObjectFit will add an Object Fit setting to the image class.
func WithImageObjectFit(fit ObjectFit) ImageOption {
	return func(i Image) Image {
		if fit > 0 {
			i.class = append(i.class, fit.String())
		}

		return i
	}
}

// NewImage creates a new image component with the given options. The created
// image can be rendered by calling its Show method.
func NewImage(options ...ImageOption) Image {
	image := Image{}

	for _, option := range options {
		image = option(image)
	}

	return image
}
