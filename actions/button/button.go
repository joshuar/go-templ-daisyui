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
	"github.com/joshuar/go-templ-daisyui/classes"
)

// Option is a functional option for a Button.
type Option components.Option[*Props]

// Props represents the properties for a DaisyUI Button component.
type Props struct {
	*attributes.Attributes
	*classes.Classes
	style     Style
	shape     Shape
	size      Size
	color     Color
	content   templ.Component
	animated  bool
	autofocus bool
	outline   bool
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
func WithSize(btnSize Size) Option {
	return func(btn *Props) {
		btn.size = btnSize
	}
}

func WithColor(color Color) Option {
	return func(btn *Props) {
		btn.color = color
	}
}

// AsShape sets the shape of the button, such as square or circle. If this option
// is not used, the button will be a regular button.
func AsShape(shape Shape) Option {
	return func(btn *Props) {
		btn.shape = shape
	}
}

// WithStyle will apply the given style to the component.
func WithStyle(value Style) Option {
	return func(btn *Props) {
		btn.style = value
	}
}

func WithName(name attributes.Name) Option {
	return func(btn *Props) {
		btn.SetName(name)
	}
}

func WithID(id string) Option {
	return func(btn *Props) {
		btn.SetID(id)
	}
}

func WithValue(value attributes.Value) Option {
	return func(btn *Props) {
		btn.SetValue(value)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(btn *Props) {
		btn.AddAttributes(attrs)
	}
}

func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		p.AddClasses(extraClasses...)
	}
}

// Build generates Button properties from the given options.
func Build(options ...Option) *Props {
	btn := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
		animated:   true,
	}

	for _, option := range options {
		option(btn)
	}

	return btn
}
