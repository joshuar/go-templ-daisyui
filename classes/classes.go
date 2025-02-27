// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package classes

import "strings"

type Class = string

type Classes struct {
	classes []Class
}

func (c *Classes) HasClasses() bool {
	return len(c.classes) > 0
}

func (c *Classes) ShowClasses() string {
	return strings.Join(c.classes, " ")
}

func (c *Classes) AddClass(class Class) {
	c.classes = append(c.classes, class)
}

// New initializes a Classes object for use in a Component.
func New() *Classes {
	return &Classes{}
}
