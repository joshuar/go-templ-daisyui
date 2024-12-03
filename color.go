// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Color -linecomment -output color_generated.go
package components

const (
	ColorNeutral          Color = iota // bg-neutral
	ColorNeutralContent                // bg-neutral-content
	ColorPrimary                       // bg-primary
	ColorPrimaryContent                // bg-primary-content
	ColorSecondary                     // bg-secondary
	ColorSecondaryContent              // bg-secondary-content
	ColorAccent                        // bg-accent
	ColorAccentContent                 // bg-accent-content
	ColorInfo                          // bg-info
	ColorInfoContent                   // bg-info-content
	ColorSuccess                       // bg-success
	ColorSuccessContent                // bg-success-content
	ColorWarning                       // bg-warning
	ColorWarningContent                // bg-warning-content
	ColorError                         // bg-error
	ColorErrorContent                  // bg-error-content
)

// Color represents a DaisyUI color.
// https://daisyui.com/docs/colors/
type Color int

// WithColor styles the component with the given Color.
func WithColor[T any](color Color) Option[T] {
	return func(c T) T {
		// Get a pointer to the underlying component.
		component := &c
		// Apply the given classes to the component.
		addClassesToComponent(component, color.String())
		// Return the modified component.
		return *component
	}
}
