// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "github.com/a-h/templ"

const (
	// Regular shaped button.
	ButtonRegular ButtonShape = iota
	// Square button with a 1:1 ratio (btn-square).
	ButtonSquare
	// Circle button with a 1:1 ratio (btn-circle).
	ButtonCircle
	// Wide button (more horizontal padding, btn-wide).
	ButtonWide
	// Full width button (btn-block).
	ButtonBlock
)

// ButtonShape sets the shape of the button.
type ButtonShape int

type modifierButtonShape struct {
	shape        ButtonShape
	shapeOutline bool
}

type buttonContent struct {
	value templ.Component
}

type ButtonProps struct {
	componentID
	componentAttributes
	modifierColor
	ModifierStateColor
	modifierResponsiveSize
	ModifierDisabled
	modifierGlass
	modifierGhost
	modifierButtonShape
	modifierAutofocus
	disableAnimation bool
	content          buttonContent
}

func (b ButtonProps) HasShape() bool {
	return b.shape > 0
}

// WithButtonShape sets the shape of the Button (square, circle, wide, block). If this
// option is not used, it will be a regular shaped Button.
func WithButtonShape(shape ButtonShape, outline bool) Option[ButtonProps] {
	return func(btn ButtonProps) ButtonProps {
		if shape != ButtonRegular {
			btn.shape = shape
			btn.shapeOutline = outline
		}

		return btn
	}
}

// WithButtonContent defines the content of the button.
func WithButtonContent(option Option[buttonContent]) Option[ButtonProps] {
	return func(btn ButtonProps) ButtonProps {
		btn.content = option(btn.content)

		return btn
	}
}

// AsTextContent will render the given text as the button content.
func AsTextContent(text string) Option[buttonContent] {
	return func(content buttonContent) buttonContent {
		content.value = PlainText(text)
		return content
	}
}

// AsIconContent will render an icon with the given options as the button content.
func AsIconContent(name string, options ...Option[IconProps]) Option[buttonContent] {
	return func(content buttonContent) buttonContent {
		content.value = Icon(name, options...)
		return content
	}
}

// DisableButtonAnimation disables the click animation on the button.
func DisableButtonAnimation() Option[ButtonProps] {
	return func(btn ButtonProps) ButtonProps {
		btn.disableAnimation = true
		return btn
	}
}

func BuildButton(options ...Option[ButtonProps]) ButtonProps {
	btn := ButtonProps{}

	for _, option := range options {
		btn = option(btn)
	}

	return btn
}
