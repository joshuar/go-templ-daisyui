// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package article

import components "github.com/joshuar/go-templ-daisyui"

const (
	Gray    ProseGrayScale = iota + 1 // prose-gray
	Slate                             // prose-slate
	Zinc                              // prose-zinc
	Neutral                           // prose-neutral
	Stone                             // prose-stone
)

type ProseGrayScale int

type Option components.Option2[*Props]

type Props struct {
	grayscale ProseGrayScale
}

// WithGrayScale adds optional gray scale scheme to the prose.
// https://github.com/tailwindlabs/tailwindcss-typography/blob/main/README.md#choosing-a-gray-scale
func WithGrayScale(scale ProseGrayScale) Option {
	return func(a *Props) {
		a.grayscale = scale
	}
}

// BuildArticle creates a new article component with the given title, content and
// options. Title is optional.
func Build(options ...Option) *Props {
	article := &Props{}

	for _, option := range options {
		option(article)
	}

	return article
}
