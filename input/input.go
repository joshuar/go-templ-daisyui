// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package input

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

const (
	AttrType = "type"
)

type Option components.Option2[*Props]

// Input is an interface to represent any type of input component.
type Input interface {
	SetValue(value attributes.Value)
	SetState(stateColor color.StateColor, outline bool)
	Show(classes ...string) templ.Component
}

type Props struct {
	FormControl *FormControlProps
	Labels      *Labels
	Input       Input
}

func (p *Props) HasLabels() bool {
	if p.Labels != nil {
		return p.Labels.left != nil || p.Labels.right != nil
	}

	return false
}

func Build(input Input, options ...Option) *Props {
	props := &Props{
		Input:       input,
		FormControl: &FormControlProps{},
		Labels:      &Labels{},
	}

	for _, option := range options {
		option(props)
	}

	return props
}
