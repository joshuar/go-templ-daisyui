// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package main

import "github.com/a-h/templ"

const (
	ListOrdered ListStyle = iota
	ListUnordered
)

type ListStyle int

type List struct {
	Attributes templ.Attributes
	Items      []templ.Component
	Style      ListStyle
}

type ListOption func(List) List

func ListAttributes(attrs templ.Attributes) ListOption {
	return func(l List) List {
		l.Attributes = attrs
		return l
	}
}

func WithItems(elements ...templ.Component) ListOption {
	return func(l List) List {
		l.Items = elements
		return l
	}
}

func NewList(style ListStyle, options ...ListOption) List {
	list := List{
		Style: style,
	}

	for _, option := range options {
		list = option(list)
	}

	return list
}
