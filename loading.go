// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

const (
	LoadingSpinner LoadingStyle = iota
	LoadingDots
	LoadingRing
	LoadingBall
	LoadingBars
	LoadingInfinity
)

type LoadingStyle int

type Loading struct {
	style LoadingStyle
	componentID
	componentAttributes
	modifierResponsiveSize
	modifierColor
}

// NewLoading creates a new loading element with the given options. The element
// can be rendered by calling the Show method.
func NewLoading(style LoadingStyle, options ...Option[Loading]) Loading {
	loading := Loading{
		style: style,
	}

	for _, option := range options {
		loading = option(loading)
	}

	return loading
}
