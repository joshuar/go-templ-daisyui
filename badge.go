// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

// Badge represents a DaisyUI badge component.
// https://daisyui.com/components/badge/
type Badge struct {
	desc string
	componentClasses
}

func (b *Badge) String() string {
	return "badge"
}

// NewBadge creates a new badge with the given options. The badge can be
// rendered by calling the Show method.
func NewBadge(desc string, options ...Option[Badge]) Badge {
	badge := Badge{
		desc: desc,
	}

	badge = WithClasses[Badge](badge.String())(badge)

	for _, option := range options {
		badge = option(badge)
	}

	return badge
}
