// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package avatar represents a DaisyUI Avatar component.
//
// https://daisyui.com/components/avatar/
package avatar

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/classes/width"
	"github.com/joshuar/go-templ-daisyui/display/image"
	"github.com/joshuar/go-templ-daisyui/layout/mask"
)

// Props are the avatar properties.
type Props struct {
	*attributes.Attributes
	*classes.Classes
	imageURL     string
	imageOptions []image.Option
	presence     Presence
}

// Option is a functional option to apply to an avatar component.
type Option components.Option[*Props]

// WithWidth sets a width class for the avatar, which can control its size.
func WithWidth(w classes.Class) Option {
	return func(p *Props) {
		p.AddClass(w)
	}
}

// WithMask option sets a mask to apply to the avatar.
func WithMask(m mask.Mask) Option {
	return func(p *Props) {
		p.AddClass(m)
	}
}

// WithPresence option sets a presence indicator on the avatar.
func WithPresence(presence Presence) Option {
	return func(p *Props) {
		p.presence = presence
	}
}

// WithImageOptions option sets additional options for the avatar image.
func WithImageOptions(options ...image.Option) Option {
	return func(p *Props) {
		p.imageOptions = options
	}
}

// WithExtraAttributes options sets any additional attributes for the component.
func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// WithExtraClasses options sets any additional classes for the component.
func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		for _, class := range extraClasses {
			p.AddClass(class)
		}
	}
}

// Build creates an avatar from the given url and options.
func Build(url string, options ...Option) *Props {
	avatar := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
		imageURL:   url,
	}

	for _, option := range options {
		option(avatar)
	}

	// Set some default class values if none specified.
	if !avatar.HasClasses() {
		WithWidth(width.W24)(avatar)
		avatar.AddClass("rounded")
	}

	return avatar
}
