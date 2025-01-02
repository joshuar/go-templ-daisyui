// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

// Badge represents a DaisyUI badge component.
// https://daisyui.com/components/badge/
type Badge struct {
	desc string
	modifierResponsiveSize
	modifierColor
	modifierState
	indicator bool
	componentClasses
}

// AsBadgeIndicator will add the "indicator" class value, when set to true, to
// allow the badge to be used as an indicator on another item.
func AsBadgeIndicator(value bool) Option[Badge] {
	return func(b Badge) Badge {
		b.indicator = value
		return b
	}
}

// WithBadgeDescription sets the badge description.
func WithBadgeDescription(desc string) Option[Badge] {
	return func(b Badge) Badge {
		b.desc = desc
		return b
	}
}

// NewBadge creates a new badge with the given options. The badge can be
// rendered by calling the Show method.
func NewBadge(options ...Option[Badge]) Badge {
	badge := Badge{}

	for _, option := range options {
		badge = option(badge)
	}

	return badge
}
