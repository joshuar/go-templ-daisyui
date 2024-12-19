// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

const (
	BadgeInfo BadgeState = iota + 1
	BadgeSuccess
	BadgeWarning
	BadgeError
)

// BadgeState defines the badge state color.
type BadgeState int

type badgeColor struct {
	color   Color
	outline bool
}

// Badge represents a DaisyUI badge component.
// https://daisyui.com/components/badge/
type Badge struct {
	desc string
	modifierSize
	state BadgeState
	color *badgeColor
}

func WithBadgeColor(color Color, outline bool) Option[Badge] {
	return func(b Badge) Badge {
		b.color = &badgeColor{
			color:   color,
			outline: outline,
		}

		return b
	}
}

// WithBadgeState sets a state color on the badge.
func WithBadgeState(state BadgeState) Option[Badge] {
	return func(b Badge) Badge {
		b.state = state
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
