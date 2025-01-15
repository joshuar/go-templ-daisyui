// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

const (
	Solid   FontAwesomeStyle = iota // fa-solid
	Regular                         // fa-regular
)

// FontAwesomeStyle is a valid FontAwesome Icon style.
type FontAwesomeStyle int

type iconLabel struct {
	label          string
	labelAlignment Alignment
}

// IconProps are the properties that can be set on an icon.
type IconProps struct {
	name    string
	style   FontAwesomeStyle
	swapon  bool
	swapoff bool
	componentTooltip
	iconLabel
}

// WithStyle assigns the given FontAwesome style to the Icon.
func WithStyle(style FontAwesomeStyle) Option[IconProps] {
	return func(icon IconProps) IconProps {
		icon.style = style
		return icon
	}
}

// WithLabel ensures the Icon has the given label.
func WithLabel(label string, alignment Alignment) Option[IconProps] {
	return func(icon IconProps) IconProps {
		icon.label = label
		icon.labelAlignment = alignment
		return icon
	}
}

// AsSwapOn allows this icon to be used as a Swap component to indicate the "on"
// state.
//
// https://daisyui.com/components/swap/
func AsSwapOn() Option[IconProps] {
	return func(icon IconProps) IconProps {
		icon.swapon = true
		return icon
	}
}

// AsSwapOff allows this icon to be used as a Swap component to indicate the
// "off" state.
//
// https://daisyui.com/components/swap/
func AsSwapOff() Option[IconProps] {
	return func(icon IconProps) IconProps {
		icon.swapoff = true
		return icon
	}
}

// FromIconProps will set an existing IconProps as the Icon properties. If
// you have previously built an Icon with BuildIcon, use this to pass the
// IconProps to Icon to render it.
func FromIconProps(props IconProps) Option[IconProps] {
	return func(_ IconProps) IconProps {
		return props
	}
}

// BuildIcon creates iconProps with the given options.
func BuildIcon(name string, options ...Option[IconProps]) IconProps {
	icon := IconProps{
		name: name,
	}

	for _, option := range options {
		icon = option(icon)
	}

	return icon
}
