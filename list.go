// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import "github.com/a-h/templ"

type list struct {
	componentAttributes
	Items []templ.Component
}

type OrderedList struct {
	list
}

type UnorderedList struct {
	list
}

func WithItems(elements ...templ.Component) Option[list] {
	return func(l list) list {
		l.Items = elements
		return l
	}
}

func NewOrderedList(options ...Option[list]) OrderedList {
	ol := OrderedList{}

	for _, option := range options {
		ol.list = option(ol.list)
	}

	return ol
}

func NewUnorderedList(options ...Option[list]) UnorderedList {
	ul := UnorderedList{}

	for _, option := range options {
		ul.list = option(ul.list)
	}

	return ul
}
