// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

type Tooltip struct {
	Tip string
	componentClasses
}

func (t *Tooltip) String() string {
	return "tooltip"
}

// Tip sets the text to display in the tooltip.
func Tip(tip string) Option[Tooltip] {
	return func(t Tooltip) Tooltip {
		t.Tip = tip
		return t
	}
}

// NewTooltop
func NewTooltip(options ...Option[Tooltip]) Tooltip {
	tooltip := Tooltip{}

	tooltip = WithClasses[Tooltip](tooltip.String())(tooltip)

	for _, option := range options {
		tooltip = option(tooltip)
	}

	return tooltip
}
