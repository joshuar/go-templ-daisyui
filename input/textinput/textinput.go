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
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Type -linecomment -output textinput.gen.go
const (
	TypeText     Type = iota // text
	TypePassword             // password
	TypeEmail                // email
	TypeURL                  // url

	AttrPlaceholder = "placeholder"
)

type Option components.Option[*Props]

// Type is the type of textinput (i.e., text, password, email, url, etc.).
type Type int

type Props struct {
	*attributes.Attributes
	size size.ResponsiveSize
	*color.ThemeColorProps
	*color.StateColorProps
	*validation.Validation
	readonly bool
}

// AsType sets the type of textinput, such as password or email. If this option
// is not used, the textinput will default to type "text".
func AsType(texttype Type) Option {
	return func(p *Props) {
		p.SetAttribute("type", texttype.String())
	}
}

// Required marks the field as required, which will trigger and enforce HTML5
// input validation.
func Required() Option {
	return func(p *Props) {
		p.SetRequired(true)
	}
}

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
		p.SetAttribute(AttrPlaceholder, text)
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
		p.ThemeColorProps = color.NewThemeColor(themeColor, false)
	}
}

func WithStateColor(stateColor color.StateColor) Option {
	return func(p *Props) {
		p.StateColorProps = color.NewStateColor(stateColor, false)
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

func Build(options ...Option) *Props {
	textinput := &Props{
		Attributes:      attributes.New(),
		Validation:      validation.New(),
		ThemeColorProps: &color.ThemeColorProps{},
		StateColorProps: &color.StateColorProps{},
	}

	for _, option := range options {
		option(textinput)
	}

	return textinput
}
