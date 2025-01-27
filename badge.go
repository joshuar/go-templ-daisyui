// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

// BadgeProps represents the properties for a Badge.
type BadgeProps struct {
	desc string
	modifierResponsiveSize
	modifierColor
	indicator bool
}

// AsBadgeIndicator will add the "indicator" class value, when set to true, to
// allow the Badge to be used as an indicator on another item.
func AsBadgeIndicator(value bool) Option[*BadgeProps] {
	return func(badge *BadgeProps) *BadgeProps {
		badge.indicator = value
		return badge
	}
}

// WithBadgeDescription sets the Badge description.
func WithBadgeDescription(desc string) Option[*BadgeProps] {
	return func(badge *BadgeProps) *BadgeProps {
		badge.desc = desc
		return badge
	}
}

// FromBadgeProps will set an existing BadgeProps as the Badge properties. If
// you have previously built a Badge with BuildBadge, use this to pass the
// BadgeProps to Badge to render it.
func FromBadgeProps(props *BadgeProps) Option[*BadgeProps] {
	return func(_ *BadgeProps) *BadgeProps {
		return props
	}
}

// BuildBadge creates a new BadgeProps with the given options.
func BuildBadge(options ...Option[*BadgeProps]) *BadgeProps {
	badge := &BadgeProps{}

	for _, option := range options {
		badge = option(badge)
	}

	return badge
}
