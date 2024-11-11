// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=InputType,InputModifier -linecomment -output input_generated.go
package components

import (
	"strings"

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
	Attributes    templ.Attributes
	Icon          *Icon
	IconAlignment Alignment
	ID            string
	Placeholder   string
	Label         string
	Error         string
	class         []string
	Type          InputType
	FormControl   bool
	Optional      bool
}

func (i Input) Class() string {
	return strings.Join(i.class, " ")
}

type InputOption func(Input) Input

func WithInputType(t InputType) InputOption {
	return func(i Input) Input {
		i.Type = t
		return i
	}
}

func WithInputModifier(m InputModifier) InputOption {
	return func(i Input) Input {
		i.class = append(i.class, m.String())
		return i
	}
}

func WithInputSize(s Size) InputOption {
	return func(i Input) Input {
		i.class = append(i.class, s.String())
		return i
	}
}

func WithPlaceholder(s string) InputOption {
	return func(i Input) Input {
		i.Placeholder = s
		return i
	}
}

func WithInputLabel(s string) InputOption {
	return func(i Input) Input {
		i.Label = s
		return i
	}
}

func WithInputIcon(icon Icon, alignment Alignment) InputOption {
	return func(i Input) Input {
		i.Icon = &icon
		i.IconAlignment = alignment
		return i
	}
}

func OptionalInput() InputOption {
	return func(i Input) Input {
		i.Optional = true
		return i
	}
}

func AsFormControl() InputOption {
	return func(i Input) Input {
		i.FormControl = true
		return i
	}
}

func WithInputAttributes(attrs templ.Attributes) InputOption {
	return func(i Input) Input {
		i.Attributes = attrs
		return i
	}
}

func NewInput(id string, options ...InputOption) Input {
	input := Input{
		ID:    id,
		class: []string{"input"},
	}

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
	input.Attributes = templ.Attributes{
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
