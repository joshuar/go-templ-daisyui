// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type list struct {
	componentAttributes
	componentItems
}

type OrderedListProps struct {
	*list
}

type UnorderedListProps struct {
	*list
}

//nolint:varnamelen
func buildOrderedList(options ...Option[*list]) *OrderedListProps {
	ol := &OrderedListProps{}

	for _, option := range options {
		ol.list = option(ol.list)
	}

	return ol
}

//nolint:varnamelen
func buildUnorderedList(options ...Option[*list]) *UnorderedListProps {
	ul := &UnorderedListProps{}

	for _, option := range options {
		ul.list = option(ul.list)
	}

	return ul
}
