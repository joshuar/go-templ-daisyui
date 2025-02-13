// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package indicator

import (
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/display/badge"
)

type Option components.Option[*BadgeIndicatorProps]

type BadgeIndicatorProps struct {
	badge *badge.Props
}

// WithBadgeOptions will set the Badge options for the Indicator Badge.
func WithBadgeOptions(options ...badge.Option) Option {
	return func(ind *BadgeIndicatorProps) {
		ind.badge = badge.Build(options...)
	}
}

// BuildBadgeIndicator creates a new Indicator as a Badge without rendering it. The properties can then
// be modified before finally rendering by calling the Show() method.
func BuildBadgeIndicator(options ...Option) *BadgeIndicatorProps {
	alert := &BadgeIndicatorProps{}

	for _, option := range options {
		option(alert)
	}

	return alert
}
