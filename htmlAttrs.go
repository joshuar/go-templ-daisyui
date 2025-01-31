// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import (
	"strconv"

	"github.com/a-h/templ"
)

// htmlAttrID is an inheritable struct for components that can have an id
// attribute.
//
// https://www.w3schools.com/TAGS/att_global_id.asp
type htmlAttrID struct {
	id   string
	name string
}

func (c *htmlAttrID) SetID(id string) {
	c.id = id
}

func (c *htmlAttrID) SetName(name string) {
	c.name = name
}

func (c *htmlAttrID) GetID() string {
	return c.id
}

func (c *htmlAttrID) GetName() string {
	return c.name
}

func (c *htmlAttrID) ID() templ.Attributes {
	attrs := make(map[string]any)
	if c.id != "" {
		attrs["id"] = c.id
	}

	if c.name != "" {
		attrs["name"] = c.name
	}

	return attrs
}

type hasID[T any] interface {
	SetID(id string)
	SetName(name string)
}

// WithID allows setting an id attribute on a component.
func WithID[T hasID[T]](id string) Option[T] {
	return func(c T) T {
		c.SetID(id)
		return c
	}
}

// WithName allows setting an name attribute on a component.
func WithName[T hasID[T]](id string) Option[T] {
	return func(c T) T {
		c.SetName(id)
		return c
	}
}

// htmlAttrValue is an inheritable struct for components that can have an value
// attribute.
type htmlAttrValue struct {
	value string
}

func (attr *htmlAttrValue) SetValue(value string) {
	attr.value = value
}

func (attr *htmlAttrValue) Value() string {
	return attr.value
}

type hasValue[T any] interface {
	SetValue(value string)
	Value() string
}

// WithValue allows setting an value on a component.
func WithValue[T hasValue[T]](value string) Option[T] {
	return func(c T) T {
		c.SetValue(value)
		return c
	}
}

// htmlAttrTabIndex is an inheritable struct for components that can have an tabindex
// attribute.
//
// https://www.w3schools.com/TAGS/att_global_tabindex.asp
type htmlAttrTabIndex struct {
	index int
}

func (attr *htmlAttrTabIndex) SetTabIndex(value int) {
	attr.index = value
}

func (attr *htmlAttrTabIndex) TabIndex() string {
	return strconv.Itoa(attr.index)
}

type hasTabIndex[T any] interface {
	SetTabIndex(value int)
	TabIndex() string
}

// WithValue allows setting an value on a component.
func WithTabIndex[T hasTabIndex[T]](value int) Option[T] {
	return func(c T) T {
		c.SetTabIndex(value)
		return c
	}
}
