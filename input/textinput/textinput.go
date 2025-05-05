// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package textinput represents a DaisyUI textinput Component.
//
// https://daisyui.com/components/input/
package textinput

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/input/validation"
)

// Option is a functional option for text input properties.
type Option components.Option[*Props]

// Props contains the properties for a text input.
type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
	*validation.Validation
	readonly bool
}

// AsType sets the type of textinput, such as password or email. If this option
// is not used, the textinput will default to type "text".
func AsType(inputType Type) Option {
	return func(p *Props) {
		p.SetAttribute("type", inputType)
	}
}

// Required marks the field as required, which will trigger and enforce HTML5
// input validation.
func Required() Option {
	return func(p *Props) {
		p.SetRequired(true)
	}
}

// ReadOnly option makes the text input readonly.
func ReadOnly() Option {
	return func(p *Props) {
		p.readonly = true
	}
}

// Validate option ensures that HTML5 validation will be applied to the input.
func Validate() Option {
	return func(p *Props) {
		p.SetValidation(true)
	}
}

func ValidateHint(hint string) Option {
	return func(p *Props) {
		p.SetHint(hint)
	}
}

func ValidatePattern(pattern string) Option {
	return func(p *Props) {
		p.SetPattern(pattern)
	}
}

func ValidateMinLength(length int) Option {
	return func(p *Props) {
		p.SetMinLength(length)
	}
}

func ValidateMaxLength(length int) Option {
	return func(p *Props) {
		p.SetMaxLength(length)
	}
}

// WithPlaceholder sets the placeholder text on the textinput.
func WithPlaceholder(text string) Option {
	return func(p *Props) {
		p.SetAttribute("placeholder", text)
	}
}

func WithSize(s Size) Option {
	return func(p *Props) {
		p.classes.Add(s)
	}
}

func WithColor(c Color) Option {
	return func(p *Props) {
		p.classes.Add(c)
	}
}

func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
	}
}

func WithID(id string) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithValue(value attributes.Value) Option {
	return func(p *Props) {
		p.SetValue(value)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func WithExtraClasses(classes ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(classes...)
	}
}

func Build(options ...Option) *Props {
	textinput := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
		Validation: validation.New(),
	}

	for _, option := range options {
		option(textinput)
	}

	return textinput
}
