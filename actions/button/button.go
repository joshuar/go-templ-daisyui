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

// WithContent sets the content for the Button.
func WithContent(content any) Option {
	return func(btn *Props) {
		btn.content = components.SetContent(content)
	}
}

// WithAutoFocus will ensure the Button is focused when displayed.
func WithAutoFocus() Option {
	return func(btn *Props) {
		btn.autofocus = true
	}
}

// WithoutClickAnimation disables the click animation on the button.
func WithoutClickAnimation() Option {
	return func(btn *Props) {
		btn.animated = false
	}
}

// WithText will set the text to display within the Badge.
func WithSize(size Size) Option {
	return func(btn *Props) {
		btn.classes.Add(size)
	}
}

func WithColor(color Color) Option {
	return func(btn *Props) {
		btn.classes.Add(color)
	}
}

// AsShape sets the shape of the button, such as square or circle. If this option
// is not used, the button will be a regular button.
func AsShape(shape Shape) Option {
	return func(btn *Props) {
		btn.classes.Add(shape)
	}
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

// WithStyle will apply the given style to the component.
func WithStyle(style Style) Option {
	return func(btn *Props) {
		btn.classes.Add(style)
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

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(btn *Props) {
		btn.attributes.AddAttributes(attrs)
	}
}

func WithExtraClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
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
