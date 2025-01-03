// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

const (
	Gray    ProseGrayScale = iota + 1 // prose-gray
	Slate                             // prose-slate
	Zinc                              // prose-zinc
	Neutral                           // prose-neutral
	Stone                             // prose-stone
)

type ProseGrayScale int

// Article represents an article component, using Tailwind's typography plugin
// and "prose" class.
// https://daisyui.com/docs/layout-and-typography/#-1
type Article struct {
	componentClasses
	grayscale ProseGrayScale
	title     string
	content   string
}

// WithGrayScale adds optional gray scale scheme to the prose.
// https://github.com/tailwindlabs/tailwindcss-typography/blob/main/README.md#choosing-a-gray-scale
func WithGrayScale(scale ProseGrayScale) Option[Article] {
	return func(a Article) Article {
		a.grayscale = scale
		return a
	}
}

// NewArticle creates a new article component with the given title, content and
// options. Title is optional.
func NewArticle(title, content string, options ...Option[Article]) Article {
	article := Article{
		title:   title,
		content: content,
	}

	for _, option := range options {
		article = option(article)
	}

	return article
}
