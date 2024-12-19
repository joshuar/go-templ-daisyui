// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

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

// cardTitle defines a Card's title.
type cardTitle struct {
	text string
	size HeaderSize
}

// Card represents a DaisyUI Card component.
// https://daisyui.com/components/card/
type Card struct {
	Body      templ.Component
	Image     *Image
	border    bool
	glass     bool
	layout    CardLayout
	fullImage bool
	centered  bool
	componentAttributes
	componentID
	componentShadow
	Actions []templ.Component
	Badges  []Badge
	title   *cardTitle
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
func WithTitle(title string, size HeaderSize) Option[Card] {
	return func(c Card) Card {
		c.title = &cardTitle{
			text: title,
			size: size,
		}

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

// WithFullImage sets the image as the card background.
func WithFullImage() Option[Card] {
	return func(c Card) Card {
		c.fullImage = true
		return c
	}
}

// WithCardLayout sets the card layout.
func WithCardLayout(l CardLayout) Option[Card] {
	return func(c Card) Card {
		c.layout = l
		return c
	}
}

// WithCenteredContent aligns any image and text content in the center of the
// card with appropriate padding.
func WithCenteredLayout() Option[Card] {
	return func(c Card) Card {
		c.centered = true
		return c
	}
}

// WithGlass makes the card glassy.
func WithGlass() Option[Card] {
	return func(c Card) Card {
		c.glass = true
		return c
	}
}

// WithBorder sets a border around the card.
func WithBorder() Option[Card] {
	return func(c Card) Card {
		c.border = true
		return c
	}
}

// WithCardShadow adds a shadow to the card.
func WithCardShadow(size Size) Option[Card] {
	return func(c Card) Card {
		c.shadowSize = size
		return c
	}
}

// NewCard creates a new card component with the given options. The card can be
// rendered by calling the Show method.
func NewCard(options ...Option[Card]) Card {
	card := Card{}

	for _, option := range options {
		card = option(card)
	}

	return card
}
