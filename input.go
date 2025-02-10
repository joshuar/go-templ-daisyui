// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type BaseInputProps struct {
	modifierResponsiveSize
	componentAttributes
	HtmlAttrID
	htmlAttrValue
	htmlAttrTabIndex
}

type modifierBordered struct {
	modifierColor
	bordered bool
}

func (m *modifierBordered) SetBordered(value bool) {
	m.bordered = value
}

// Bordered will return true if the Component should have a bordered class.
func (m *modifierBordered) Bordered() bool {
	return m.bordered
}

type customBordered[T any] interface {
	SetBordered(value bool)
}

// Bordered will set whether the input should have a border.
func Bordered[T customBordered[T]](value bool) Option[T] {
	return func(c T) T {
		c.SetBordered(value)
		return c
	}
}

type modifierPlaceholder struct {
	placeholderText string
}

func (m *modifierPlaceholder) SetPlaceholder(text string) {
	m.placeholderText = text
}

// PlaceholderIsSet will return true if the Component has placeholder text.
func (m modifierPlaceholder) PlaceholderIsSet() bool {
	return m.placeholderText != ""
}

type customPlaceholder[T any] interface {
	SetPlaceholder(text string)
}

// WithPlaceholder will set the placeholder text.
func WithPlaceholder[T customPlaceholder[T]](text string) Option[T] {
	return func(c T) T {
		c.SetPlaceholder(text)
		return c
	}
}

type modifierTextValidation struct {
	required bool
	min      int
	max      int
	pattern  string
}

func (m modifierTextValidation) IsRequired() bool {
	return m.required
}

func (m *modifierTextValidation) SetValidation() {
	m.required = true
}

type customValidation[T any] interface {
	IsRequired() bool
	SetValidation()
}

// WithPlaceholder will set the placeholder text.
func WithValidationRequired[T customValidation[T]]() Option[T] {
	return func(c T) T {
		c.SetValidation()
		return c
	}
}
