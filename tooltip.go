// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "strings"

type Tooltip struct {
	Tip   string
	class []string
}

func (t Tooltip) Class() string {
	return strings.Join(t.class, " ")
}

type TooltipOption Option[Tooltip]

// Tip sets the text to display in the tooltip.
func Tip(tip string) TooltipOption {
	return func(t Tooltip) Tooltip {
		t.Tip = tip
		return t
	}
}

// TooltipAlignment sets where the tooltip will open. Valid values of Alignment
// are Top, Bottom, Left or Right. Other Alignment values, while applied, may
// not render corrrectly.
func ToolTipAlignment(a Alignment) TooltipOption {
	return func(t Tooltip) Tooltip {
		t.class = append(t.class, "tooltip-"+a.String())
		return t
	}
}

// TooltipColor sets the color of the tooltip.
func ToolTipColor(c Color) TooltipOption {
	return func(t Tooltip) Tooltip {
		t.class = append(t.class, c.String())
		return t
	}
}

func NewTooltip(options ...TooltipOption) Tooltip {
	tooltip := Tooltip{
		class: []string{"tooltip"},
	}

	for _, option := range options {
		tooltip = option(tooltip)
	}

	return tooltip
}
