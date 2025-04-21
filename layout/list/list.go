// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package list

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
)

type Option components.Option[*List]

type List struct {
	*components.Props
}

func WithID(id string) Option {
	return func(l *List) {
		l.SetID(id)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(l *List) {
		l.AddAttributes(attrs)
	}
}

type OrderedListProps struct {
	list *List
}

type UnorderedListProps struct {
	list *List
}

func BuildOrdered(options ...Option) *OrderedListProps {
	ol := &OrderedListProps{
		list: &List{
			Props: components.InitProps(),
		},
	}

	for _, option := range options {
		option(ol.list)
	}

	return ol
}

func BuildUnordered(options ...Option) *UnorderedListProps {
	ul := &UnorderedListProps{
		list: &List{
			Props: components.InitProps(),
		},
	}

	for _, option := range options {
		option(ul.list)
	}

	return ul
}
