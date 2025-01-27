// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

const (
	Z0    ZIndex = iota + 1 // z-index: 0;
	Z10                     // z-index: 10;
	Z20                     // z-index: 20;
	Z30                     // z-index: 30;
	Z40                     // z-index: 40;
	Z50                     // z-index: 50;
	ZAuto                   // z-index: auto;
)

type ZIndex int

// modifierZIndex can be embedded into components to allow setting a z-index option
// on the component. The component will need to handle rendering with the
// appropriate size itself.
type modifierZIndex struct {
	zindex   ZIndex
	negative bool
}

func (m *modifierZIndex) SetZIndex(index ZIndex, negative bool) {
	m.zindex = index
	m.negative = negative
}

func (m *modifierZIndex) ZIndexIsSet() bool {
	return m.zindex > 0
}

type customZIndex[T any] interface {
	SetZIndex(index ZIndex, negative bool)
}

// WithZIndex adds the given z-index to the component.
func WithZIndex[T customZIndex[T]](index ZIndex, negative bool) Option[T] {
	return func(c T) T {
		c.SetZIndex(index, negative)
		return c
	}
}
