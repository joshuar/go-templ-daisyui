// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package placeholder

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/display/avatar"
	"github.com/joshuar/go-templ-daisyui/layout/mask"
)

const (
	MD Size = iota
	LG
	XXL
	SM
)

type Size int

// PlaceholderProps are the properties for an avatar placeholder.
type Props struct {
	*attributes.Attributes
	*classes.Classes
	size     Size
	value    string
	presence avatar.Presence
}

// Option is a functional option to apply to an avatar component.
type Option components.Option[*Props]

// WithMask option sets a mask to apply to the avatar.
func WithMask(m mask.Mask) Option {
	return func(p *Props) {
		p.AddClass(m)
	}
}

// WithPresence option sets a presence indicator on the avatar.
func WithPresence(presence avatar.Presence) Option {
	return func(p *Props) {
		p.presence = presence
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
func Build(value string, size Size, options ...Option) *Props {
	avatar := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
		value:      value,
		size:       size,
	}

	for _, option := range options {
		option(avatar)
	}

	avatar.applyClasses()

	return avatar
}
