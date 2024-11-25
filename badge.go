// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "strings"

// Badge represents a DaisyUI badge component.
// https://daisyui.com/components/badge/
type Badge struct {
	desc    string
	classes []string
}

func (b Badge) Class() string {
	return strings.Join(b.classes, " ")
}

// BadgeOption adds an option to a card.
type BadgeOption Option[Badge]

// BadgeSize adds the given size to the badge class.
func BadgeSize(size Size) BadgeOption {
	return func(b Badge) Badge {
		b.classes = append(b.classes, size.String())
		return b
	}
}

// BadgeColor adds the given color to the badge class.
func BadgeColor(color Color) BadgeOption {
	return func(b Badge) Badge {
		b.classes = append(b.classes, color.String())
		return b
	}
}

// NewBadge creates a new badge with the given options. The badge can be
// rendered by calling the Show method.
func NewBadge(desc string, options ...BadgeOption) Badge {
	badge := Badge{
		desc:    desc,
		classes: []string{"badge"},
	}

	for _, option := range options {
		badge = option(badge)
	}

	return badge
}
