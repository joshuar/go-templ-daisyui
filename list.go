// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type List struct {
	componentAttributes
	componentItems
}

type OrderedListProps struct {
	*List
}

type UnorderedListProps struct {
	*List
}

//nolint:varnamelen
func buildOrderedList(options ...Option[*List]) *OrderedListProps {
	ol := &OrderedListProps{}

	for _, option := range options {
		ol.List = option(ol.List)
	}

	return ol
}

//nolint:varnamelen
func buildUnorderedList(options ...Option[*List]) *UnorderedListProps {
	ul := &UnorderedListProps{}

	for _, option := range options {
		ul.List = option(ul.List)
	}

	return ul
}
