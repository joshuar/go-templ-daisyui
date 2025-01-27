// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

type ToggleProps struct {
	BaseInputProps
	modifierDisabled
	modifierChecked
	modifierColor
	htmlAttrID
}

func BuildToggle(options ...Option[*ToggleProps]) *ToggleProps {
	toggle := &ToggleProps{}

	for _, option := range options {
		toggle = option(toggle)
	}

	return toggle
}
