// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=ButtonModifier,ButtonShape -linecomment -output button_generated.go
package main

import (
	"strings"

	"github.com/a-h/templ"
)

const (
	ButtonNoModifier ButtonModifier = iota //
	ButtonNeutral                          // btn-neutral
	ButtonPrimary                          // btn-primary
	ButtonSecondary                        // btn-secondary
	ButtonAccent                           // btn-accent
	ButtonInfo                             // btn-info
	ButtonSuccess                          // btn-success
	ButtonWarning                          // btn-warning
	ButtonError                            // btn-error
	ButtonLink                             // btn-link
	ButtonGhost                            // btn-ghost
)

type ButtonModifier int

const (
	ButtonRegular ButtonShape = iota //
	ButtonSquare                     // btn-square
	ButtonCircle                     // btn-circle
)

type ButtonShape int

type Button struct {
	Icon       *IconLayout
	Attributes templ.Attributes
	Label      string
	ID         string
	class      []string
	Shape      ButtonShape
}

func (b Button) Class() string {
	return strings.Join(b.class, " ")
}

type ButtonOption func(Button) Button

func WithModifier(m ButtonModifier) ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, m.String())
		return b
	}
}

func WithSize(s Size) ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, s.String())
		return b
	}
}

func Active() ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, "btn-active")
		return b
	}
}

func Outline() ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, "btn-outline")
		return b
	}
}

func Wide() ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, "btn-wide")
		return b
	}
}

func Glass() ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, "glass")
		return b
	}
}

func Disabled() ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, "btn-disabled")
		return b
	}
}

func WithIcon(i Icon, l Alignment) ButtonOption {
	return func(b Button) Button {
		b.Icon = &IconLayout{Icon: i, Alignment: l}
		return b
	}
}

func AsShape(s ButtonShape) ButtonOption {
	return func(b Button) Button {
		b.Shape = s
		return b
	}
}

func WithButtonClasses(classes ...string) ButtonOption {
	return func(b Button) Button {
		b.class = append(b.class, classes...)
		return b
	}
}

func ButtonAttributes(attrs templ.Attributes) ButtonOption {
	return func(b Button) Button {
		b.Attributes = attrs
		return b
	}
}

func NewButton(label, id string, options ...ButtonOption) Button {
	button := Button{
		Label: label,
		ID:    id,
		class: []string{"btn"},
	}

	for _, option := range options {
		button = option(button)
	}

	return button
}
