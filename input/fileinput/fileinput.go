// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package fileinput

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/input/validation"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

type Props struct {
	*attributes.Attributes
	*classes.Classes
	*validation.Validation
	themeColor color.ThemeColor
	stateColor color.StateColor
	size       size.ResponsiveSize
}

type Option components.Option[*Props]

// WithSize option sets the responsive size of the component.
func WithSize(inputSize size.ResponsiveSize) Option {
	return func(input *Props) {
		input.size = inputSize
	}
}

// WithThemeColor option sets a theme color on the component.
func WithThemeColor(themeColor color.ThemeColor) Option {
	return func(input *Props) {
		input.themeColor = themeColor
	}
}

// WithStateColor option sets a state color on the component.
func WithStateColor(stateColor color.StateColor) Option {
	return func(input *Props) {
		input.stateColor = stateColor
	}
}

// WithName option sets a name attribute on the component.
func WithName(name attributes.Name) Option {
	return func(input *Props) {
		input.SetName(name)
	}
}

// WithID option sets an id attribute on the component.
func WithID(id string) Option {
	return func(input *Props) {
		input.SetID(id)
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

// WithExtraAttributes option sets additional attributes on the component.
func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(input *Props) {
		input.AddAttributes(attrs)
	}
}

// WithExtraClasses option assigns additional classes to the component.
func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(input *Props) {
		input.AddClasses(extraClasses...)
	}
}

func Build(options ...Option) *Props {
	props := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
		Validation: validation.New(),
	}

	for _, option := range options {
		option(props)
	}

	return props
}
