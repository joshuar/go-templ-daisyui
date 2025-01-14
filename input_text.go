// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"github.com/a-h/templ"
)

const (
	TextInputText TextInputType = iota
	TextInputPassword
	TextInputEmail
	TextInputURL
)

type TextInputType int

type textInputInsideLabels struct {
	labelLeft  templ.Component
	labelRight templ.Component
}

type textInputOutsideLabels struct {
	labelTopLeft     string
	labelTopRight    string
	labelBottomLeft  string
	labelBottomRight string
}

// SetTopLeftLabel will set the top-left outside label on the Text Input.
func (l *textInputOutsideLabels) SetTopLeftLabel(label string) {
	l.labelTopLeft = label
}

// SetTopRightLabel will set the top-right outside label on the Text Input.
func (l *textInputOutsideLabels) SetTopRightLabel(label string) {
	l.labelTopRight = label
}

// SetBottomLeftLabel will set the bottom-left outside label on the Text Input.
func (l *textInputOutsideLabels) SetBottomLeftLabel(label string) {
	l.labelBottomLeft = label
}

// SetBottomRightLabel will set the bottom-right outside label on the Text Input.
func (l *textInputOutsideLabels) SetBottomRightLabel(label string) {
	l.labelBottomRight = label
}

func (l *textInputOutsideLabels) OutsideLabelsSet() bool {
	return l.labelTopLeft != "" || l.labelTopRight != "" || l.labelBottomLeft != "" || l.labelBottomRight != ""
}

type TextInputProps struct {
	BaseInputProps
	modifierBordered
	modifierPlaceholder
	modifierGhost
	modifierTextValidation
	inputType    TextInputType
	InsideLabels *textInputInsideLabels
	textInputOutsideLabels
	formcontrol bool
	indicator   string
}

// WithFormControl will ensure the Input Text Component is wrapped with a
// "form-control" class.
func WithFormControl() Option[TextInputProps] {
	return func(input TextInputProps) TextInputProps {
		input.formcontrol = true
		return input
	}
}

// WithInsideLabels assigns the given components as the left and right labels,
// inside the Text Input box. Both labels are optional.
func WithInsideLabels(left, right templ.Component) Option[TextInputProps] {
	return func(input TextInputProps) TextInputProps {
		labels := &textInputInsideLabels{}

		if left != nil {
			labels.labelLeft = left
		}

		if right != nil {
			labels.labelRight = right
		}

		input.InsideLabels = labels

		return input
	}
}

// WithOutsideLabels assigns the given components as the top/bottom left/right labels,
// around the Text Input box. All labels are optional.
func WithOutsideLabels(topLeft, topRight, bottomLeft, bottomRight string) Option[TextInputProps] {
	return func(input TextInputProps) TextInputProps {
		if topLeft != "" {
			input.labelTopLeft = topLeft
		}

		if topRight != "" {
			input.labelTopRight = topRight
		}

		if bottomLeft != "" {
			input.labelBottomLeft = bottomLeft
		}

		if bottomRight != "" {
			input.labelBottomRight = bottomRight
		}

		return input
	}
}

// AsPassword indicates that this TextInput component should be treated as a
// password and its input obscured.
func AsPassword() Option[TextInputProps] {
	return func(input TextInputProps) TextInputProps {
		input.inputType = TextInputPassword
		return input
	}
}

// AsEmail indicates that this TextInput component should be treated as a
// email.
func AsEmail() Option[TextInputProps] {
	return func(input TextInputProps) TextInputProps {
		input.inputType = TextInputEmail
		return input
	}
}

// AsEmail indicates that this TextInput component should be treated as a
// URL.
func AsURL() Option[TextInputProps] {
	return func(input TextInputProps) TextInputProps {
		input.inputType = TextInputURL
		return input
	}
}

// FromTextInputProps will set an existing TextInputProps as the TextInput properties. If
// you have previously built a TextInput with BuildTextInput, use this to pass the
// TextInputProps to TextInput to render it.
func FromTextInputProps(props TextInputProps) Option[TextInputProps] {
	return func(_ TextInputProps) TextInputProps {
		return props
	}
}

func BuildTextInput(options ...Option[TextInputProps]) TextInputProps {
	props := TextInputProps{}

	for _, option := range options {
		props = option(props)
	}

	return props
}
