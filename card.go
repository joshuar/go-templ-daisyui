// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

//go:generate go run golang.org/x/tools/cmd/stringer -type=CardLayout -linecomment -output card_generated.go
package components

import (
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
	Body  templ.Component
	Image *Image
	componentAttributes
	componentID
	componentClasses
	Actions []templ.Component
	Badges  []Badge
	Title   Header
}

func (c *Card) String() string {
	return "card"
}

// WithBody sets the container for card content.
func WithBody(t templ.Component) Option[Card] {
	return func(c Card) Card {
		c.Body = t
		return c
	}
}

// WithActions sets the container for card actions.
func WithActions(actions ...templ.Component) Option[Card] {
	return func(c Card) Card {
		c.Actions = append(c.Actions, actions...)
		return c
	}
}

// WithBadges adds the provided list of badges to the Card.
func WithBadges(badges ...Badge) Option[Card] {
	return func(c Card) Card {
		c.Badges = append(c.Badges, badges...)
		return c
	}
}

// WithTitle sets the title for card body.
func WithTitle(title string) Option[Card] {
	return func(c Card) Card {
		c.Title = NewHeader(title,
			WithHeaderSize(H2),
			WithClasses[Header]("card-title"))

		return c
	}
}

// WithImage sets the figure displayed in the card body.
func WithImage(i Image) Option[Card] {
	return func(c Card) Card {
		c.Image = &i
		return c
	}
}

// WithCardLayout sets the card layout.
func WithCardLayout(l CardLayout) Option[Card] {
	return func(c Card) Card {
		c.classes = append(c.classes, l.String())
		return c
	}
}

// WithGlass makes the card glassy.
func WithGlass() Option[Card] {
	return func(c Card) Card {
		c.classes = append(c.classes, "glass")
		return c
	}
}

// WithBorder sets a border around the card.
func WithBorder() Option[Card] {
	return func(c Card) Card {
		c.classes = append(c.classes, "card-bordered")
		return c
	}
}

// WithFullImage sets the image as the card background.
func WithFullImage() Option[Card] {
	return func(c Card) Card {
		c.classes = append(c.classes, "image-full")
		return c
	}
}

// WithCardShadow adds a shadow to the card.
func WithCardShadow(size Size) Option[Card] {
	return func(c Card) Card {
		c.classes = append(c.classes, "shadow-"+size.String())
		return c
	}
}

// NewCard creates a new card component with the given options. The card can be
// rendered by calling the Show method.
func NewCard(options ...Option[Card]) Card {
	card := Card{}

	card = WithClasses[Card](card.String())(card)

	for _, option := range options {
		card = option(card)
	}

	return card
}
