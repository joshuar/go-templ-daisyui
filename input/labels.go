// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package input

import (
	"github.com/a-h/templ"
	"github.com/joshuar/go-templ-daisyui/display/icon"
)

const (
	Left  LabelPosition = iota // Left label.
	Right                      // Right label.
)

// LabelPosition defines the position of a label on an input.
type LabelPosition int

type Labels struct {
	left  templ.Component
	right templ.Component
}

func (l *Labels) SetLabel(position LabelPosition, label any) {
	switch position {
	case Left:
		l.left = newLabel(label)
	case Right:
		l.right = newLabel(label)
	}
}

func WithLabel(position LabelPosition, label any) Option {
	return func(p *Props) {
		p.Labels.SetLabel(position, label)
	}
}

// newLabel converts a label of any valid type into a templ.Component. Invalid
// types are silently ignored and the function will return nil.
func newLabel(label any) templ.Component {
	switch value := label.(type) {
	case string:
		return templ.Raw(value)
	case icon.Props:
		return value.Show()
	case templ.Component:
		return value
	default:
		return nil
	}
}
