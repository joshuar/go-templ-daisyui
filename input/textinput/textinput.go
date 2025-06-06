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
	validation *validation.Validation
}

// AsType sets the type of textinput, such as password or email. If this option
// is not used, the textinput will default to type "text".
func AsType(inputType Type) Option {
	return func(p *Props) {
		p.attributes.SetAttribute("type", string(inputType))
	}
}

// Required marks the field as required, which will trigger and enforce HTML5
// input validation.
func Required() Option {
	return func(p *Props) {
		p.validation.SetRequired(true)
	}
}

// ReadOnly option makes the text input readonly.
func ReadOnly() Option {
	return func(p *Props) {
		p.attributes.SetAttribute("readonly", true)
	}
}

// Validate option ensures that HTML5 validation will be applied to the input.
func Validate() Option {
	return func(p *Props) {
		p.validation.SetValidation(true)
	}
}

func ValidateHint(hint string) Option {
	return func(p *Props) {
		p.validation.SetHint(hint)
	}
}

func ValidatePattern(pattern string) Option {
	return func(p *Props) {
		p.validation.SetPattern(pattern)
	}
}

func ValidateMinLength(length int) Option {
	return func(p *Props) {
		p.validation.SetMinLength(length)
	}
}

func ValidateMaxLength(length int) Option {
	return func(p *Props) {
		p.validation.SetMaxLength(length)
	}
}

// WithPlaceholder sets the placeholder text on the textinput.
func WithPlaceholder(text string) Option {
	return func(p *Props) {
		p.validation.SetAttribute("placeholder", text)
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
		p.attributes.SetName(name)
	}
}

func WithID(id string) Option {
	return func(p *Props) {
		p.attributes.SetID(id)
	}
}

func WithValue(value attributes.Value) Option {
	return func(p *Props) {
		p.attributes.SetValue(value)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
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
		validation: validation.New(),
	}

	for _, option := range options {
		option(textinput)
	}

	return textinput
}

func (i *Props) GetID() string {
	return i.attributes.GetID()
}

func (i *Props) SetValue(value attributes.Value) {
	i.attributes.SetValue(value)
}
