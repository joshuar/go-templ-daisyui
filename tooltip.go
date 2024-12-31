// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

type Tooltip struct {
	Tip string
	componentClasses
}

type componentTooltip struct {
	componentAttributes
	componentClasses
}

// customisableAttributes is an interface to detect components whose attributes
// can be customized.
type customisableTooltip interface {
	SetAttribute(key string, value any)
	AddClasses(classes ...string)
}

func WithToolTip[T any](tip string) Option[T] {
	return func(c T) T {
		// Get a pointer to the underlying component.
		component := &c
		// If the component supports customizing attributes, set the attributes
		// to the given value.
		if settable, ok := any(component).(customisableTooltip); ok {
			settable.SetAttribute("data-tip", tip)
			settable.AddClasses("tooltip")
		}
		// Return the modified component.
		return *component
	}
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

// NewTooltip creates a new Tooltip component with the specified options.
func NewTooltip(options ...Option[Tooltip]) Tooltip {
	tooltip := Tooltip{}

	tooltip = WithClasses[Tooltip](tooltip.String())(tooltip)

	for _, option := range options {
		tooltip = option(tooltip)
	}

	return tooltip
}
