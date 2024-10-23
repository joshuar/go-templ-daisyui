// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package main

import (
	"github.com/a-h/templ"
)

type Form struct {
	ID         string
	Inputs     *OrderedMap[string, Input]
	Attributes templ.Attributes
	Buttons    *OrderedMap[string, Button]
	Info       string
}

func (f Form) GetInputs() []Input {
	var inputs []Input

	iterator := f.Inputs.Iterator()

	for {
		i, _, v := iterator()
		if i == nil {
			break
		}

		inputs = append(inputs, v)
	}

	return inputs
}

func (f Form) GetButtons() []Button {
	var buttons []Button

	iterator := f.Buttons.Iterator()

	for {
		i, _, v := iterator()
		if i == nil {
			break
		}

		buttons = append(buttons, v)
	}

	return buttons
}

type FormOption func(Form) Form

func Info(info string) FormOption {
	return func(f Form) Form {
		f.Info = info
		return f
	}
}

func FormAttributes(attrs templ.Attributes) FormOption {
	return func(f Form) Form {
		f.Attributes = attrs
		return f
	}
}

func Buttons(buttons ...Button) FormOption {
	return func(f Form) Form {
		for _, button := range buttons {
			f.Buttons.Set(button.ID, button)
		}

		return f
	}
}

func Inputs(inputs ...Input) FormOption {
	return func(f Form) Form {
		for _, input := range inputs {
			f.Inputs.Set(input.ID, input)
		}

		return f
	}
}

func NewForm(id string, options ...FormOption) Form {
	form := Form{
		ID:      id,
		Inputs:  NewOrderedMap[string, Input](),
		Buttons: NewOrderedMap[string, Button](),
	}

	for _, option := range options {
		form = option(form)
	}

	return form
}
