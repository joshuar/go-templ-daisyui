// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

const (
	XS ResponsiveSize = iota + 1 // Extra-small size. class: {component}-xs
	SM                           // Small size. class: {component}-sm
	MD                           // Medium size. class: {component}-md
	LG                           // Large size. class: {component}-lg
	XL                           // Extra-large size. class: {component}-xl
)

// ResponsizeSize represents one of the responsive size values (e.g., "md",
// "lg", etc.).
type ResponsiveSize int

// modifierResponsiveSize can be embedded into components to allow setting a size option
// on the component. The component will need to handle rendering with the
// appropriate size itself.
type modifierResponsiveSize struct {
	size ResponsiveSize
}

func (m modifierResponsiveSize) SetSize(size ResponsiveSize) modifierResponsiveSize {
	m.size = size
	return m
}

// SizeIsSet will return true if the Component has a size.
func (m *modifierResponsiveSize) SizeIsSet() bool {
	return m.size > 0
}

type hasResponsiveSizeModifier[T any] interface {
	SetSize(size ResponsiveSize) T
}

// WithSize adds the given size to the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithResponsiveSize[T any](size ResponsiveSize) Option[T] {
	return func(c T) T {
		if settable, ok := any(c).(hasResponsiveSizeModifier[T]); ok {
			return settable.SetSize(size)
		}
		slog.Warn("WithResponsiveSize passed as option to component, but component does not support responsive size modifier.",
			slog.String("component", fmt.Sprintf("%T", &c)))

		return c
	}
}

const (
	Size0 Size = iota + 1
	SizePx
	Size0_5
	Size1
	Size1_5
	Size2
	Size2_5
	Size3
	Size3_5
	Size4
	Size5
	Size6
	Size7
	Size8
	Size9
	Size10
	Size11
	Size12
	Size14
	Size16
	Size20
	Size24
	Size28
	Size32
	Size36
	Size40
	Size44
	Size48
	Size52
	Size56
	Size64
	Size72
	Size80
	Size96
	SizeAuto
	SizeOneHalf
	SizeOneThird
	SizeTwoThirds
	SizeOneQuarter
	SizeTwoQuarters
	SizeThreeQuarters
	SizeOneFifth
	SizeTwoFifths
	SizeThreeFifths
	SizeFourFifths
	SizeOneSixth
	SizeTwoSixths
	SizeThreeSixths
	SizeFourSixths
	SizeFiveSixths
	SizeMin
	SizeMax
)

// Size represents a fixed, percentage or one of auto/min/max sizes.
//
// https://tailwindcss.com/docs/size
type Size int

type modifierSize struct {
	size Size
}

func (m *modifierSize) SetSize(size Size) {
	m.size = size
}

// SizeIsSet will return true if the Component has a size.
func (m *modifierSize) SizeIsSet() bool {
	return m.size > 0
}

type customisableSize interface {
	SetSize(size Size)
}

// WithSize adds the given size to the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithSize[T any](size Size) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableSize); ok {
			settable.SetSize(size)
		} else {
			slog.Warn("WithSize passed as option to component, but component does not support size modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}

// WithSizeAuto sets `size-auto` on the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithSizeAuto[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableSize); ok {
			settable.SetSize(SizeAuto)
		} else {
			slog.Warn("WithSize passed as option to component, but component does not support size modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}

// WithSizeMin sets `size-min` on the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithSizeMin[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableSize); ok {
			settable.SetSize(SizeMin)
		} else {
			slog.Warn("WithSize passed as option to component, but component does not support size modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}

// WithSizeMax sets `size-max` on the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithSizeMax[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableSize); ok {
			settable.SetSize(SizeMax)
		} else {
			slog.Warn("WithSize passed as option to component, but component does not support size modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
