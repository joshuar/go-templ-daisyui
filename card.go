// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

//go:generate go run golang.org/x/tools/cmd/stringer -type=CardLayout -linecomment -output card_generated.go
package components

import (
	"strings"

	"github.com/a-h/templ"
)

const (
	CardLayoutNormal         CardLayout = iota // card-normal
	CardLayoutCompact                          // card-compact
	CardLayoutSide                             // card-side
	CardLayoytSideResponsive                   // lg:card-side
)

// CardLayout is used to define the card style.
type CardLayout int

// Card represents a DaisyUI Card component.
// https://daisyui.com/components/card/
type Card struct {
	Body    templ.Component
	Actions templ.Component
	Image   *Image
	Title   string
	ID      string
	class   []string
	Layout  CardLayout
}

func (c Card) Class() string {
	return strings.Join(c.class, " ")
}

// CardOption adds an option to a card.
type CardOption Option[Card]

// WithBody sets the container for card content.
func WithBody(t templ.Component) CardOption {
	return func(c Card) Card {
		c.Body = t
		return c
	}
}

// WithActions sets the container for card actions.
func WithActions(t templ.Component) CardOption {
	return func(c Card) Card {
		c.Actions = t
		return c
	}
}

// WithTitle sets the title for card body.
func WithTitle(t string) CardOption {
	return func(c Card) Card {
		c.Title = t
		return c
	}
}

// WithImage sets the figure displayed in the card body.
func WithImage(i Image) CardOption {
	return func(c Card) Card {
		c.Image = &i
		return c
	}
}

// WithCardLayout sets the card layout.
func WithCardLayout(l CardLayout) CardOption {
	return func(c Card) Card {
		c.class = append(c.class, l.String())
		return c
	}
}

// WithGlass makes the card glassy.
func WithGlass() CardOption {
	return func(c Card) Card {
		c.class = append(c.class, "glass")
		return c
	}
}

// WithBorder sets a border around the card.
func WithBorder() CardOption {
	return func(c Card) Card {
		c.class = append(c.class, "card-bordered")
		return c
	}
}

// WithFullImage sets the image as the card background.
func WithFullImage() CardOption {
	return func(c Card) Card {
		c.class = append(c.class, "image-full")
		return c
	}
}

// NewCard creates a new card component with the given options. The card can be
// rendered by calling the Show method.
func NewCard(id string, options ...CardOption) Card {
	input := Card{
		ID:    id,
		class: []string{"card"},
	}

	for _, option := range options {
		input = option(input)
	}

	return input
}
