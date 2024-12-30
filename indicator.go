// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "github.com/a-h/templ"

const (
	AlignIndicatorStart IndicatorAlignment = iota
	AlignIndicatorCenter
	AlignIndicatorEnd
	AlignIndicatorTop
	AlignIndicatorMiddle
	AlignIndicatorBottom
)

type IndicatorAlignment int

type Indicator struct {
	alignment IndicatorAlignment
}

type CardIndicator struct {
	Indicator
	ItemContent templ.Component
	Card        templ.Component
}

// WithBody sets the container for card content.
func WithInidicatorAlignment(a IndicatorAlignment) Option[Indicator] {
	return func(ind Indicator) Indicator {
		ind.alignment = a
		return ind
	}
}

func NewCardIndicator(align IndicatorAlignment, item, card templ.Component) CardIndicator {
	ind := CardIndicator{}
	ind.alignment = align
	ind.ItemContent = item
	ind.Card = card

	return ind
}
