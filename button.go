// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=ButtonModifier,ButtonShape -linecomment -output button_generated.go
package components

const (
	ButtonNoModifier ButtonModifier = iota //
	ButtonNeutral                          // btn-neutral
	ButtonPrimary                          // btn-primary
	ButtonSecondary                        // btn-secondary
	ButtonAccent                           // btn-accent
	ButtonInfo                             // btn-info
	ButtonSuccess                          // btn-success
	ButtonWarning                          // btn-warning
	ButtonError                            // btn-error
	ButtonLink                             // btn-link
	ButtonGhost                            // btn-ghost
)

type ButtonModifier int

const (
	ButtonRegular ButtonShape = iota //
	ButtonSquare                     // btn-square
	ButtonCircle                     // btn-circle
)

type ButtonShape int

type Button struct {
	Icon *Icon
	componentAttributes
	Tooltip *Tooltip
	Label   string
	ID      string
	componentClasses
	IconAlignment Alignment
}

func (b *Button) String() string {
	return "btn"
}

func WithModifier(m ButtonModifier) Option[Button] {
	return func(b Button) Button {
		b.AddClasses(m.String())
		return b
	}
}

func Active() Option[Button] {
	return func(b Button) Button {
		b.AddClasses("btn-active")
		return b
	}
}

func Outline() Option[Button] {
	return func(b Button) Button {
		b.AddClasses("btn-outline")
		return b
	}
}

func Wide() Option[Button] {
	return func(b Button) Button {
		b.AddClasses("btn-wide")
		return b
	}
}

func Glass() Option[Button] {
	return func(b Button) Button {
		b.AddClasses("glass")
		return b
	}
}

func Disabled() Option[Button] {
	return func(b Button) Button {
		b.AddClasses("btn-disabled")
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
			ToolTipAlignment(alignment),
			Tip(tip),
		)
		btn.Tooltip = &tooltip

		return btn
	}
}

func AsShape(shape ButtonShape) Option[Button] {
	return func(b Button) Button {
		if shape != ButtonRegular {
			b.AddClasses(shape.String())
		}

		return b
	}
}

func NewButton(label, id string, options ...Option[Button]) Button {
	btn := Button{
		Label: label,
		ID:    id,
	}

	btn = WithClasses[Button](btn.String())(btn)

	for _, option := range options {
		btn = option(btn)
	}

	return btn
}
