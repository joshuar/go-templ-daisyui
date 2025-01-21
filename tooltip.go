// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

type Tooltip struct {
	Tip string
}

type componentTooltip struct {
	componentAttributes
}

// customisableAttributes is an interface to detect components whose attributes
// can be customized.
type customTooltip[T any] interface {
	SetAttribute(key string, value any)
	AddClasses(classes ...string)
}

func WithToolTip[T customTooltip[T]](tip string) Option[T] {
	return func(c T) T {
		c.SetAttribute("data-tip", tip)
		c.AddClasses("tooltip")

		return c
	}
}

func (t *Tooltip) String() string {
	return "tooltip"
}

// Tip sets the text to display in the tooltip.
func Tip(tip string) Option[*Tooltip] {
	return func(t *Tooltip) *Tooltip {
		t.Tip = tip
		return t
	}
}

// NewTooltip creates a new Tooltip component with the specified options.
func NewTooltip(options ...Option[*Tooltip]) *Tooltip {
	tooltip := &Tooltip{}

	for _, option := range options {
		tooltip = option(tooltip)
	}

	return tooltip
}
