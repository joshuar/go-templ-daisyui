// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

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

const (
	// Place the Button in the top left of the container.
	// Adds classes: absolute left-2 top-2.
	TopLeft ButtonContainerAlignment = iota + 1
	// Place the Button in the top right of the container.
	// Adds classes: absolute right-2 top-2.
	TopRight
)

// ButtonContainerAlignment defines the Button placement within its container
// element.
type ButtonContainerAlignment int

// Button represents a DaisyUI Button component.
// https://daisyui.com/components/button/
type Button struct {
	Icon *Icon
	componentAttributes
	componentClasses
	Tooltip            *Tooltip
	Label              string
	ID                 string
	IconAlignment      Alignment
	containerAlignment ButtonContainerAlignment
	shape              ButtonShape
	active             bool
	disabled           bool
	modifierColor
	modifierGlass
	modifierResponsiveSize
	modifierState
}

// IsActive returns whether the Button has the active modifier.
func (b *Button) IsActive() bool {
	return b.active
}

// SetActive sets the active modifier on an existing Button. This can be used
// after creation of a Button to toggle the active state.
func (b *Button) SetActive(value bool) {
	b.active = value
}

// IsDisabled returns whether the Button has the disabled modifier.
func (b *Button) IsDisabled() bool {
	return b.disabled
}

// SetDisabled sets the disabled modifier on an existing Button. This can be
// used after the creation of a Button to toggle the active state.
func (b *Button) SetDisabled(value bool) {
	b.disabled = value
}

// Active ensures the Button has the active modifier.
func Active() Option[Button] {
	return func(b Button) Button {
		b.active = true
		return b
	}
}

// Disabled ensures the Button has the disabled modifier.
func Disabled() Option[Button] {
	return func(b Button) Button {
		b.disabled = true
		return b
	}
}

// WithButtonShape sets the shape of the Button (square, circle, wide, block). If this
// option is not used, it will be a regular shaped Button.
func WithButtonShape(shape ButtonShape) Option[Button] {
	return func(b Button) Button {
		if shape != ButtonRegular {
			b.shape = shape
		}

		return b
	}
}

func WithIcon(icon Icon, alignment Alignment) Option[Button] {
	return func(b Button) Button {
		b.Icon = &icon
		b.IconAlignment = alignment

		return b
	}
}

// WithButtonTooltip adds a tooltip to the button.
func WithButtonTooltip(tip string, alignment Alignment) Option[Button] {
	return func(btn Button) Button {
		tooltip := NewTooltip(
			WithAlignment[Tooltip](alignment),
			Tip(tip),
		)
		btn.Tooltip = &tooltip

		return btn
	}
}

// WithContainerAlignment positions the button within its container element.
func WithContainerAlignment(alignment ButtonContainerAlignment) Option[Button] {
	return func(btn Button) Button {
		btn.containerAlignment = alignment
		return btn
	}
}

func NewButton(label, id string, options ...Option[Button]) Button {
	btn := Button{
		Label: label,
		ID:    id,
	}

	for _, option := range options {
		btn = option(btn)
	}

	return btn
}
