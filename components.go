// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import (
	"github.com/a-h/templ"
)

// Option represents a generic option that can be applied to a component.
type Option[T any] func(T) T

type Option2[T any] func(T)

// Component represents any DaisyUI component.
type Component interface {
	Show(classes ...string) templ.Component
}

// setContent converts the content of any valid type into a templ.Component to
// be used within a Component as its content. Invalid types are silently ignored
// and the function will return nil.
func SetContent(content any) templ.Component {
	switch value := content.(type) {
	case string:
		return templ.Raw(value)
	case Component:
		return value.Show()
	case templ.Component:
		return value
	default:
		return nil
	}
}

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
