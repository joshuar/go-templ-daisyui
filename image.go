// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

// ImageProps represents the properties for an image.
type ImageProps struct {
	url  string
	alt  string
	lazy bool
	modifierMask
	modifierSize
	modifierObjectFit
}

// WithLazyLoading will add the `loading="lazy"` attribute to the image.
func WithLazyLoading() Option[*ImageProps] {
	return func(image *ImageProps) *ImageProps {
		image.lazy = true
		return image
	}
}

// WithAltText sets the alt text to be displayed for the image.
func WithAltText(alt string) Option[*ImageProps] {
	return func(image *ImageProps) *ImageProps {
		image.alt = alt
		return image
	}
}

// FromImageProps will set an existing ImageProps as the Image properties. If
// you have previously built an Image with BuildImage, use this to pass the
// ImageProps to Image to render it.
func FromImageProps(props *ImageProps) Option[*ImageProps] {
	return func(_ *ImageProps) *ImageProps {
		return props
	}
}

// BuildImage creates a new ImageProps.
func BuildImage(url string, options ...Option[*ImageProps]) *ImageProps {
	image := &ImageProps{
		url: url,
	}

	for _, option := range options {
		image = option(image)
	}

	return image
}
