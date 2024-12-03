// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=InputType,InputModifier -linecomment -output input_generated.go
package components

import (
	"github.com/a-h/templ"
)

const (
	InputTypeText     InputType = iota // text
	InputTypePassword                  // password
	InputTypeSearch                    // search
)

type InputType int

const (
	InputNoModifier InputModifier = iota //
	InputBordered                        // input-bordered
	InputGhost                           // input-ghost
	InputPrimary                         // input-primary
	InputSecondary                       // input-secondary
	InputAccent                          // input-accent
	InputInfo                            // input-info
	InputSuccess                         // input-success
	InputWarning                         // input-warning
	InputError                           // input-error
)

type InputModifier int

type Input struct {
	componentAttributes
	componentClasses
	Icon          *Icon
	ID            string
	Placeholder   string
	Label         string
	Error         string
	Value         string
	IconAlignment Alignment
	Type          InputType
	FormControl   bool
	Optional      bool
}

func (i *Input) String() string {
	return ""
}

// SetValue assigns the given value to the input.
func (i *Input) SetValue(value string) {
	i.Value = value
}

// WithInputType defines the type of input (i.e., text, password, search, etc.).
func WithInputType(t InputType) Option[Input] {
	return func(i Input) Input {
		i.Type = t
		return i
	}
}

// WithInputModifier applies the given DaisyUI modifier to the input.
func WithInputModifier(m InputModifier) Option[Input] {
	return func(i Input) Input {
		i.AddClasses(m.String())
		return i
	}
}

// WithPlaceholder assigns placeholder text for when the input has no current
// value.
func WithPlaceholder(s string) Option[Input] {
	return func(i Input) Input {
		i.Placeholder = s
		return i
	}
}

// WithInputLabel assigns a label to the input.
func WithInputLabel(s string) Option[Input] {
	return func(i Input) Input {
		i.Label = s
		return i
	}
}

// WithInputIcon assigns an icon to the input. The icon can be aligned either
// Left or Right.
func WithInputIcon(icon Icon, alignment Alignment) Option[Input] {
	return func(i Input) Input {
		i.Icon = &icon
		i.IconAlignment = alignment

		return i
	}
}

// OptionalInput if set, will mark the input with a badge indicating it is
// optional.
func OptionalInput() Option[Input] {
	return func(i Input) Input {
		i.Optional = true
		return i
	}
}

// AsFormControl adds necessary attributes to the input to allow it to be used
// within form elements.
func AsFormControl() Option[Input] {
	return func(i Input) Input {
		i.FormControl = true
		return i
	}
}

// NewInput creates a new input with the given values.
func NewInput(id string, options ...Option[Input]) Input {
	input := Input{
		ID: id,
	}

	input = WithClasses[Input](input.String())(input)

	for _, option := range options {
		input = option(input)
	}

	return input
}

type SearchInput struct {
	Validator   string
	Value       string
	LoadingText string
	Input
}

func (si *SearchInput) GenerateInput() Input {
	input := si.Input
	input.attributes = templ.Attributes{
		"name":         "Terms",
		"value":        si.Value,
		"tabindex":     0,
		"hx-post":      si.Validator,
		"hx-trigger":   "input changed delay:500ms, SearchRequest",
		"hx-target":    "#search-results",
		"hx-indicator": ".htmx-indicator",
	}

	return input
}
