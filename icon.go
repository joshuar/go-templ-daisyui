// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

const (
	Solid   FontAwesomeStyle = iota // fa-solid
	Regular                         // fa-regular
)

// FontAwesomeStyle is a valid FontAwesome Icon style.
type FontAwesomeStyle int

// IconProps are the properties that can be set on an icon.
type IconProps struct {
	name    string
	label   string
	style   FontAwesomeStyle
	swapon  bool
	swapoff bool
	componentTooltip
}

// WithStyle assigns the given FontAwesome style to the Icon.
func WithStyle(s FontAwesomeStyle) Option[IconProps] {
	return func(i IconProps) IconProps {
		i.style = s
		return i
	}
}

// WithLabel ensures the Icon has the given label.
func WithLabel(l string) Option[IconProps] {
	return func(i IconProps) IconProps {
		i.label = l
		return i
	}
}

// AsSwapOn allows this icon to be used as a Swap component to indicate the "on"
// state.
//
// https://daisyui.com/components/swap/
func AsSwapOn() Option[IconProps] {
	return func(i IconProps) IconProps {
		i.swapon = true
		return i
	}
}

// AsSwapOff allows this icon to be used as a Swap component to indicate the
// "off" state.
//
// https://daisyui.com/components/swap/
func AsSwapOff() Option[IconProps] {
	return func(i IconProps) IconProps {
		i.swapoff = true
		return i
	}
}

// FromIconProps will set an existing IconProps as the Icon properties. If
// you have previously built an Icon with BuildIcon, use this to pass the
// IconProps to Icon to render it.
func FromIconProps(props IconProps) Option[IconProps] {
	return func(ip IconProps) IconProps {
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
