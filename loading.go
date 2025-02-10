// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

const (
	LoadingSpinner LoadingStyle = iota
	LoadingDots
	LoadingRing
	LoadingBall
	LoadingBars
	LoadingInfinity
)

// LoadingStyle represents the DaisyUI loading style.
//
// https://daisyui.com/components/loading/
type LoadingStyle int

// LoadingProps are the properties of a Loading component.
type LoadingProps struct {
	style         LoadingStyle
	htmxIndicator bool
	HtmlAttrID
	componentAttributes
	modifierResponsiveSize
	modifierColor
}

// AsHTMXIndicator adds the "htmx-indicator" class to the Loading component,
// which will allow HTMX to add styling and control to the component.
func AsHTMXIndicator() Option[*LoadingProps] {
	return func(lp *LoadingProps) *LoadingProps {
		lp.htmxIndicator = true
		return lp
	}
}

// FromLoadingProps will set an existing LoadingProps as the Loading properties. If
// you have previously built an LoadingProps with BuildLoading, use this to pass the
// LoadingProps to Loading to render it.
func FromLoadingProps(props *LoadingProps) Option[*LoadingProps] {
	return func(_ *LoadingProps) *LoadingProps {
		return props
	}
}

// BuildLoading creates a new LoadingProps.
func BuildLoading(style LoadingStyle, options ...Option[*LoadingProps]) *LoadingProps {
	loading := &LoadingProps{
		style: style,
	}

	for _, option := range options {
		loading = option(loading)
	}

	return loading
}
