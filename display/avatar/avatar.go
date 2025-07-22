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
)

// Props are the avatar properties.
type Props struct {
	attributes  *attributes.Attributes
	classes     *components.Classes
	presence    components.Class
	placeholder bool
}

// Option is a functional option to apply to an avatar component.
type Option components.Option[*Props]

// AsPlaceholder option sets the avatar-placeholder class.
func AsPlaceholder() Option {
	return func(p *Props) {
		p.placeholder = true
	}
}

// WithPresence option sets a presence indicator on the avatar.
func WithPresence(presence Presence) Option {
	return func(p *Props) {
		p.presence = presence
	}
}

// WithAttributes options sets any additional attributes for the component.
func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

// WithClasses options sets any additional classes for the component.
func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// Build creates an avatar from the given url and options.
func Build(options ...Option) *Props {
	avatar := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(avatar)
	}

	return avatar
}
