// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import (
	"log/slog"

	"github.com/a-h/templ"
)

// Option represents a generic option that can be applied to a component.
type Option[T any] func(T) T

type Option2[T any] func(T)

// componentItems is an inheritable struct for components that have "children"
// items, that could be any other component.
type componentItems struct {
	items []templ.Component
}

func (c *componentItems) SetItems(items ...templ.Component) {
	c.items = append(c.items, items...)
}

type customItems[T any] interface {
	SetItems(items ...templ.Component)
}

// WithValue allows setting an value on a component.
func WithItems[T customItems[T]](items ...templ.Component) Option[T] {
	return func(c T) T {
		c.SetItems(items...)
		return c
	}
}

const (
	TopLeft ContentLocation = iota
	TopRight
	BottomLeft
	BottomRight
)

type ContentLocation int

type borderContent struct {
	TopRight    templ.Component
	BottomRight templ.Component
	TopLeft     templ.Component
	BottomLeft  templ.Component
}

func (c *borderContent) SetBorderContent(location ContentLocation, content templ.Component) {
	switch location {
	case TopRight:
		c.TopRight = content
	case TopLeft:
		c.TopLeft = content
	case BottomRight:
		c.BottomRight = content
	case BottomLeft:
		c.BottomLeft = content
	default:
		slog.Debug("Unsupported location for border content.",
			slog.Any("location", location))
	}
}

type hasBorderContent[T any] interface {
	SetBorderContent(location ContentLocation, content templ.Component)
}

// WithValue allows setting an value on a component.
func WithBorderContent[T hasBorderContent[T]](location ContentLocation, content templ.Component) Option[T] {
	return func(c T) T {
		c.SetBorderContent(location, content)
		return c
	}
}
