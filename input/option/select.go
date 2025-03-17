// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package option is a DaisyUI select input component.
//
// https://daisyui.com/components/select/
package option

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/input/validation"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

type Props struct {
	*validation.Validation
	*attributes.Attributes
	label      string
	size       size.ResponsiveSize
	themeColor color.ThemeColor
	stateColor color.StateColor
	values     []Value
}

// Value represents an option value in a select. It can be disabled or selected.
type Value struct {
	Label    string
	Value    string
	Disabled bool
	Selected bool
}

type Option components.Option[*Props]

// WithLabel sets a "label" for the select. This is implemented as a disabled and selected by
// default option.
func WithLabel(label string) Option {
	return func(p *Props) {
		p.label = label
	}
}

// Required marks the field as required, which will trigger and enforce HTML5
// input validation.
func Required() Option {
	return func(p *Props) {
		p.SetRequired(true)
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

// WithText will set the text to display within the Badge.
func WithSize(inputSize size.ResponsiveSize) Option {
	return func(p *Props) {
		p.size = inputSize
	}
}

func WithThemeColor(themeColor color.ThemeColor) Option {
	return func(p *Props) {
		p.themeColor = themeColor
	}
}

func WithStateColor(stateColor color.StateColor) Option {
	return func(p *Props) {
		p.stateColor = stateColor
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

func WithValues(values ...Value) Option {
	return func(p *Props) {
		p.values = values
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func Build(options ...Option) *Props {
	textinput := &Props{
		Attributes: attributes.New(),
		Validation: validation.New(),
	}

	for _, option := range options {
		option(textinput)
	}

	return textinput
}
