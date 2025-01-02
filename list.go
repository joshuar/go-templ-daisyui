// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type list struct {
	componentAttributes
	componentItems
}

type OrderedList struct {
	list
}

type UnorderedList struct {
	list
}

//nolint:varnamelen
func NewOrderedList(options ...Option[list]) OrderedList {
	ol := OrderedList{}

	for _, option := range options {
		ol.list = option(ol.list)
	}

	return ol
}

//nolint:varnamelen
func NewUnorderedList(options ...Option[list]) UnorderedList {
	ul := UnorderedList{}

	for _, option := range options {
		ul.list = option(ul.list)
	}

	return ul
}
