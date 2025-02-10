// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package button

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/display/icon"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

const (
	// Square button with a 1:1 ratio (btn-square).
	Square Shape = iota + 1
	// Circle button with a 1:1 ratio (btn-circle).
	Circle
	// Wide button (more horizontal padding, btn-wide).
	Wide
	// Full width button (btn-block).
	Block
)

type Option components.Option2[*Props]

// Shape sets the shape of the button.
type Shape int

type modifierShape struct {
	shape        Shape
	shapeOutline bool
}

type Props struct {
	*attributes.Attributes
	size size.ResponsiveSize
	color.Colors
	*modifierShape
	content          templ.Component
	noClickAnimation bool
	autofocus        bool
}

// AsShape sets the shape of the button, such as square or circle. If this option
// is not used, the button will be a regular button.
func AsShape(shape Shape, outline bool) Option {
	return func(p *Props) {
		if shape == 0 {
			return
		}

		p.shape = shape
		p.shapeOutline = outline
	}
}

// WithContent sets the content for the Button.
func WithContent(content any) Option {
	return func(p *Props) {
		p.content = setContent(content)
	}
}

// WithAutoFocus will ensure the Button is focused when displayed.
func WithAutoFocus() Option {
	return func(p *Props) {
		p.autofocus = true
	}
}

// WithoutClickAnimation disables the click animation on the button.
func WithoutClickAnimation() Option {
	return func(p *Props) {
		p.noClickAnimation = true
	}
}

// WithText will set the text to display within the Badge.
func WithSize(btnSize size.ResponsiveSize) Option {
	return func(p *Props) {
		p.size = btnSize
	}
}

func WithThemeColor(themeColor color.ThemeColor, outline bool) Option {
	return func(p *Props) {
		p.SetColor(themeColor, outline)
	}
}

func WithStateColor(stateColor color.StateColor, outline bool) Option {
	return func(p *Props) {
		p.SetState(stateColor, outline)
	}
}

func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
	}
}

func WithID(id attributes.ID) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithValue(value attributes.Value) Option {
	return func(p *Props) {
		p.SetValue(value)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

func Build(options ...Option) *Props {
	btn := &Props{
		Attributes:    attributes.New(),
		modifierShape: &modifierShape{},
	}

	for _, option := range options {
		option(btn)
	}

	return btn
}

// setContent converts the content of any valid type into a templ.Component to
// be used as the Button content. Invalid types are silently ignored and the
// function will return nil.
func setContent(content any) templ.Component {
	switch value := content.(type) {
	case string:
		return templ.Raw(value)
	case *icon.Props:
		return value.Show()
	case templ.Component:
		return value
	default:
		return nil
	}
}
