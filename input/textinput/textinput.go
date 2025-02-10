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

type Option components.Option2[*Props]

// Type is the type of textinput (i.e., text, password, email, url, etc.).
type Type int

type Props struct {
	*attributes.Attributes
	size size.ResponsiveSize
	color.Colors
	bordered bool
	required bool
	readonly bool
}

// AsType sets the type of textinput, such as password or email. If this option
// is not used, the textinput will default to type "text".
func AsType(texttype Type) Option {
	return func(p *Props) {
		p.SetAttribute("type", texttype.String())
	}
}

// Bordered will ensure the textinput has a visible border around it.
func Bordered() Option {
	return func(p *Props) {
		p.bordered = true
	}
}

// Required marks the field as required, which will trigger and enforce HTML5
// input validation.
func Required() Option {
	return func(p *Props) {
		p.required = true
	}
}

func ReadOnly() Option {
	return func(p *Props) {
		p.readonly = true
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

func WithThemeColor(themeColor color.ThemeColor, outline bool) Option {
	return func(p *Props) {
		p.SetColor(themeColor, outline)
	}
}

func WithStateColor(stateColor color.StateColor, outline bool) Option {
	return func(p *Props) {
		p.SetState(stateColor, outline)
	}
}

func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
	}
}

func WithID(id attributes.ID) Option {
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
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(textinput)
	}

	return textinput
}
