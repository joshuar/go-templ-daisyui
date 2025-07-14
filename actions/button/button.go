// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package button represents a DaisyUI Button component.
//
// https://daisyui.com/components/button/
package button

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

// Option is a functional option for a Button.
type Option components.Option[*Props]

// Props represents the properties for a DaisyUI Button component.
type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
	content    templ.Component
	animated   bool
	autofocus  bool
	outline    bool
}

func (p *Props) GetID() string {
	return p.attributes.GetID()
}

func (p *Props) SetValue(value attributes.Value) {
	p.attributes.SetValue(value)
}

func AsSubmit() Option {
	return func(p *Props) {
		p.attributes.SetAttribute("type", "submit")
	}
}

func AsReset() Option {
	return func(p *Props) {
		p.attributes.SetAttribute("type", "reset")
	}
}

func WithName(name attributes.Name) Option {
	return func(btn *Props) {
		btn.attributes.SetName(name)
	}
}

func WithID(id string) Option {
	return func(btn *Props) {
		btn.attributes.SetID(id)
	}
}

func WithValue(value attributes.Value) Option {
	return func(btn *Props) {
		btn.attributes.SetValue(value)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(btn *Props) {
		btn.attributes.AddAttributes(attrs)
	}
}

func WithClasses(classes ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(classes...)
	}
}

// Build generates Button properties from the given options.
func Build(options ...Option) *Props {
	btn := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
		animated:   true,
	}

	for _, option := range options {
		option(btn)
	}

	return btn
}
